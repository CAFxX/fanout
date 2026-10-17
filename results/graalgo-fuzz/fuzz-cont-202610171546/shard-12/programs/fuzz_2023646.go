package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 bool = false
    v2 := []int{}
    var v3 int32 = -545
    var v4 uint8 = 255
    var v5 int16 = 0
    var v6 float32 = (float32(-2.25) + -2.25)
    var v7 int16 = (v5 >> ((4294967295) & 63))
    var v8 int = 32767
    v9 := GenS1{F0: int8(v5)}
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: 42}
        v5 = 318
        _ = fi10
        _ = v11
    }
    if (v1 || (!(0 < 1))) {
        var v12 string = "foo bar"
        v13 := GenS2{F0: v8, F1: (v4 % (((v4 + v4)) | 1))}
        if v1 {
            if v1 {
                for fi14 := 0; fi14 < 10; fi14++ {
                    _ = len(v2)
                    var v15 float32 = 1e10
                    _ = fi14
                    _ = v15
                }
                var v16 int8 = 0
                _ = v16
            } else {
                for fi17 := 0; fi17 < 13; fi17++ {
                    if len(v2) > 0 {
                        v2[0] = v8
                    }
                    v18 := GenS2{F0: -2147483648, F1: 255}
                    var v19 int64 = 9223372036854775807
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            if v1 {
                switch v12 {
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "foo bar":
                    v6 += 3.14159
                case "hello":
                    var v20 uint = 422
                    _ = v20
                }
                if v1 {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        v22 := GenS3{F0: (v8 + -502), F1: 523}
                        _ = fi21
                        _ = v22
                    }
                } else {
                    if true {
                        if (173 <= v4) {
                            var v23 uint64 = 0
                            _ = v23
                        }
                        if false {
                            var v24 uint16 = 1
                            if (v6 < ((v6 - v6) - (3.14159 * ((v6 + (v6 + v6)) - v6)))) {
                                v2 = append(v2, -251)
                            }
                            fmt.Println(v12, v4, v7)
                            _ = v24
                        } else {
                            for fi25 := 0; fi25 < 2; fi25++ {
                                v26 := GenS4{F0: 2}
                                v5 -= v7
                                _ = fi25
                                _ = v26
                            }
                            var v27 uint = 18446744073709551615
                            _ = v27
                        }
                        v8 -= (v8 + v8)
                    } else {
                        v1 = v1
                    }
                    v1 = (42 > 0)
                    v28 := GenS3{F0: v8, F1: 1}
                    _ = v28
                }
            }
        } else {
            var v29 int64 = 893
            var v30 string = "foo bar"
            _ = v29
            _ = v30
        }
        _ = v12
        _ = v13
    } else {
        if v1 {
            v5 = v5
            var v31 uint8 = v4
            _ = v31
        } else {
            if false {
                var v32 float64 = 1.5
                v6 = 1e10
                _ = v32
            }
        }
    }
    var v33 uint16 = 0
    var v34 int64 = int64(v3)
    v33 += (^(v33))
    v33 = (^(v33))
    v1 = ((-(v8)) < 127)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
