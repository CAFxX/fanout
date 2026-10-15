package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 int64 }
func genFn0() string {
    var v7 int16 = 32767
    v7 = (v7 * v7)
    _ = v7
    return "hello"
}
func genFn1() int16 {
    if (((454 == 781) || false) && ((!false) || (42 <= 1000))) {
        var v8 uint8 = 255
        for fi9 := 0; fi9 < 12; fi9++ {
            v8 -= v8
            _ = fi9
        }
        v10 := GenS1{F0: -65, F1: 255, F2: 99}
        _ = v10
        _ = v8
    }
    v11 := GenS1{F0: 100, F1: 32767, F2: 1}
    v12 := []int{-398}
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 uint32 = 965
    var v2 uint64 = 909
    var v3 float64 = (0.0 / 3.14159)
    var v4 uint = 2
    var v5 int16 = 1000
    v6 := make(map[int]bool)
    v3 = float64(0.0)
    switch ("\n\t" + ("a" + "hello")) {
    case "":
        var v13 int16 = v5
        _ = v13
    case "foo bar":
        delete(v6, 128)
    }
    var v14 uint16 = 65535
    fmt.Println(v1, v2)
    v6[-419] = true
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            v3 += v3
            v17 := GenS1{F0: 100, F1: v14, F2: 127}
            if (("\n\t" + "\n\t") > "\n\t") {
                v18 := v6[(fi16 ^ fi15)]
                v19 := GenS1{F0: -66, F1: 552, F2: 0}
                var v20 string = "a"
                v5 -= (v5 * (v5 - (v5 * -1)))
                v4 += v4
                _ = v18
                _ = v19
                _ = v20
            }
            v21 := GenS2{F0: fi16, F1: 246}
            if true {
                switch -129 {
                case -1:
                    switch "hello" {
                    case "foo bar":
                        var v22 int32 = int32(fi15)
                        _ = v22
                    case "\n\t":
                        if false {
                            v23 := v6[32768]
                            var v24 float32 = 0.1
                            fmt.Println(v23, v5)
                            _ = v23
                            _ = v24
                        } else {
                            var v25 int = (fi15 | -912)
                            if false {
                                var v26 uint = (v4 + (^(v4)))
                                _ = v26
                            }
                            _ = v25
                        }
                    default:
                        if (int32(v2) >= 0) {
                            v2 += uint64(v1)
                            if (true || true) {
                                if (!true) {
                                    if (v5 < (v5 & (v5 * (-(v5))))) {
                                        switch "" {
                                        case "a":
                                            var v27 int64 = 0
                                            _ = v27
                                        case "":
                                            v28 := GenS1{F0: 100, F1: v14, F2: 1}
                                            _ = v28
                                        default:
                                            v14 += uint16(v4)
                                        }
                                        v29 := GenS2{F0: (fi15 * fi16), F1: 222}
                                        _ = v29
                                    }
                                } else {
                                    switch -708 {
                                    case -2147483648:
                                        v30 := v6[32767]
                                        v3 += v3
                                        if v30 {
                                            v31 := []int{fi16}
                                            v2 += 0
                                            v31 = append(v31, int(v5))
                                            v32 := GenS1{F0: 100, F1: v14, F2: 42}
                                            fmt.Println(v30, v30, v30)
                                            if (-1 <= -378) {
                                                if v30 {
                                                    if false {
                                                        v33 := GenS1{F0: 90, F1: v14, F2: 26}
                                                        _ = v33
                                                    }
                                                    var v34 int8 = 1
                                                    delete(v6, (fi16 | -9223372036854775808))
                                                    _ = v34
                                                }
                                                fmt.Println(fi15, v4, v30)
                                                v2 -= v2
                                            } else {
                                                var v35 bool = (v2 <= uint64(v1))
                                                _ = v35
                                            }
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = v30
                                    case 128:
                                        if false {
                                            var v36 float32 = 0.0
                                            var v37 bool = true
                                            _ = v36
                                            _ = v37
                                        } else {
                                            if false {
                                                switch (fi16 | -368) {
                                                case 4294967295:
                                                    v6[(fi15 >> ((v4) & 63))] = false
                                                }
                                                v38 := []int{fi16, fi16, 435, fi16}
                                                v6[fi15] = (((fi16 % ((fi15) | 1)) <= (fi15 ^ (fi15 >> ((v4) & 63)))) || false)
                                                delete(v6, fi16)
                                                _ = v38
                                            } else {
                                                fmt.Println(v2, v2)
                                                v39 := v6[(fi16 - fi15)]
                                                var v40 uint64 = (v2 + uint64(v14))
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                    }
                                    v1 = (v1 - (v1 + v1))
                                }
                                var v41 int32 = 42
                                _ = v41
                            }
                            _ = len(v6)
                        } else {
                            var v42 int = -9223372036854775808
                            _ = v42
                        }
                    }
                case -758:
                    var v43 int32 = 421
                    _ = v43
                }
            }
            var v44 bool = false
            v6[0] = (42 >= 128)
            _ = fi16
            _ = v17
            _ = v21
            _ = v44
        }
        if (!(!(!true))) {
            var v45 int = (fi15 << ((v4) & 63))
            var v46 float32 = 1e10
            v4 -= 65535
            _ = v45
            _ = v46
        } else {
            for fi47 := 0; fi47 < 11; fi47++ {
                fmt.Println(v14, v14, v14)
                switch "a" {
                case "":
                    v2 -= (v2 ^ (v2 >> ((2) & 63)))
                }
                v48 := []int{}
                _ = fi47
                _ = v48
            }
            v14 = 65535
        }
        _ = fi15
    }
    var v49 float64 = float64((((float32((v3 + -2.25)) / ((0.1 + (0.0 - (0.1 + (float32(3.14159) - float32(float64(-2.25)))))) + 3.14159)) - float32(float64(-2.25))) * float32(v3)))
    if ((v14 << ((v4) & 63)) != (v14 + (v14 >> (((v4 >> ((4294967295) & 63))) & 63)))) {
        fmt.Println(v2)
        for fi50 := 0; fi50 < 17; fi50++ {
            if (false || ((1.5 - 3.14159) > 0.1)) {
                v51 := GenS3{F0: (v4 + 0), F1: fi50, F2: 0}
                v3 -= (v3 + -2.25)
                v2 += 1
                _ = v51
            }
            _ = fi50
        }
        var v52 bool = (!((v5 >> ((707) & 63)) <= (v5 | (v5 * (^(v5))))))
        _ = v52
    } else {
        if (127 == 42) {
            v3 += 3.14159
        }
        v5 = (v5 - v5)
    }
    var v53 int = 810
    v6[186] = (4294967295 > v53)
    for fi54 := 0; fi54 < 10; fi54++ {
        if (!(false && (3.14159 != 3.14159))) {
            var v55 string = (("" + "\n\t") + "")
            var v56 float32 = -2.25
            _ = v55
            _ = v56
        }
        _ = fi54
    }
    var v57 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v53
    _ = v57
    _ = v6
}
