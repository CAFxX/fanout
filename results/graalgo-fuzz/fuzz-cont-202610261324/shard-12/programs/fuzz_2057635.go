package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint8; F1 uint64 }
func genFn0() bool {
    v6 := []int{-38}
    if (42 > 515) {
        var v7 uint8 = 1
        if true {
            switch 128 {
            case -1:
                v8 := GenS1{F0: 42}
                _ = v8
            default:
                if (1e10 == (float64(0.1) / float64((1.5 * -2.25)))) {
                    v7 -= v7
                }
            }
            v9 := GenS2{F0: -340, F1: 0}
            _ = v9
        }
        _ = v7
    } else {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 uint8 = 127
            var v12 string = "hello"
            v13 := GenS1{F0: v11}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        var v14 string = ""
        _ = v14
    }
    v6 = append(v6, 32767)
    _ = v6
    return true
}

func main() {
    v1 := []uint32{65535, 1, 2147483648}
    v2 := []uint32{374, 554, 758}
    var v3 uint8 = 128
    var v4 uint32 = 2
    v5 := []uint32{(v4 / ((v4) | 1)), (v4 + (v4 - v4))}
    var v15 string = ("\n\t" + "a")
    var v16 float32 = float32((0.1 - (float64((3.14159 / (3.14159 / (3.14159 - 1e10)))) / 1e10)))
    if true {
        var v17 int16 = 42
        _ = v17
    } else {
        var v18 uint8 = v3
        v19 := GenS1{F0: 2}
        switch (v15 + ("hello" + v15)) {
        case "\n\t":
            v15 = v15
        case "foo bar":
            var v20 string = (v15 + "foo bar")
            _ = v20
        }
        _ = v18
        _ = v19
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 8; fi22++ {
            if (!true) {
                v4 = v4
                switch v15 {
                case "a":
                    var v23 uint8 = 128
                    _ = v23
                default:
                    if (true || ((-9223372036854775808 > fi22) || (!(true && (79 <= 621))))) {
                        _ = len(v2)
                    }
                }
                if (float32((1.5 / (-2.25 - 1e10))) > (v16 * v16)) {
                    v24 := GenS3{F0: 32767, F1: 1000}
                    v4 = 2147483648
                    v1 = append(v1, 65535)
                    _ = v24
                } else {
                    if (42 != -2147483648) {
                        v15 = (v15 + "")
                    }
                }
            } else {
                fmt.Println(v4)
                switch ((v15 + "hello") + v15) {
                case "foo bar":
                    v25 := GenS1{F0: 128}
                    _ = v25
                case "\n\t":
                    var v26 uint8 = (v3 >> ((235) & 63))
                    _ = v26
                default:
                    if (255 <= 2) {
                        v5 = append(v5, (v4 + (v4 << ((18446744073709551615) & 63))))
                        v27 := GenS1{F0: (^(v3))}
                        v3 = 255
                        _ = v27
                    } else {
                        v28 := GenS2{F0: 238, F1: -32768}
                        _ = v28
                    }
                }
                if (281 < 42) {
                    v16 += 1.5
                    var v29 float64 = 3.14159
                    _ = v29
                } else {
                    var v30 int64 = 181
                    _ = v30
                }
            }
            var v31 string = v15
            if len(v2) > 0 {
                v2[0] = v4
            }
            v15 = (("\n\t" + (v15 + "")) + v31)
            fmt.Println(fi21, v16, v4)
            _ = fi22
            _ = v31
        }
        if true {
            v32 := GenS2{F0: fi21, F1: -538}
            if false {
                for fi33 := 0; fi33 < 4; fi33++ {
                    if true {
                        v34 := GenS4{F0: (v3 & (v3 - (v3 | 1))), F1: 460, F2: 127}
                        if len(v1) > 0 {
                            v1[0] = v4
                        }
                        _ = v34
                    } else {
                        if len(v5) > 0 {
                            v5[0] = v4
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        switch fi33 {
                        case 1:
                            var v35 bool = ((fi33 <= fi33) || true)
                            _ = v35
                        }
                    }
                    v16 = 0.1
                    _ = fi33
                }
                if false {
                    if false {
                        switch fi21 {
                        case 894:
                            for fi36 := 0; fi36 < 16; fi36++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch fi36 {
                                case 128:
                                    var v37 int = -9223372036854775808
                                    _ = v37
                                }
                                switch (^(fi36)) {
                                case 4294967295:
                                    var v38 int = fi36
                                    _ = v38
                                case 9223372036854775807:
                                    var v39 int16 = 42
                                    _ = v39
                                default:
                                    switch fi21 {
                                    case 127:
                                        if (!false) {
                                            fmt.Println(v15, v16)
                                        }
                                    }
                                    switch (fi21 | fi21) {
                                    case -2147483648:
                                        v3 -= v3
                                    case 0:
                                        fmt.Println(fi36)
                                    case 403:
                                        var v40 uint = 42
                                        v4 += 149
                                        var v41 uint32 = (v4 * uint32(v3))
                                        _ = v40
                                        _ = v41
                                    }
                                    var v42 int32 = -1
                                    _ = v42
                                }
                                _ = fi36
                            }
                        case -241:
                            v43 := GenS3{F0: 80, F1: 42}
                            _ = v43
                        }
                    } else {
                        v44 := GenS1{F0: v3}
                        var v45 int8 = -128
                        _ = v44
                        _ = v45
                    }
                    _ = len(v1)
                } else {
                    v46 := GenS5{F0: 0, F1: 1, F2: (^(v3))}
                    _ = v46
                }
            }
            _ = v32
        } else {
            if (65535 >= 892) {
                var v47 uint32 = v4
                if (42 > -1) {
                    _ = len(v5)
                }
                _ = v47
            }
        }
        var v48 float64 = 0.0
        _ = fi21
        _ = v48
    }
    v49 := GenS6{F0: 42, F1: 4294967295}
    var v50 uint = 18446744073709551615
    var v51 int = -2147483648
    v52 := GenS3{F0: 255, F1: -32768}
    v53 := GenS1{F0: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
