package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int64; F1 int64; F2 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int }
func genFn0(p0 bool, p1 uint16, p2 int) uint64 {
    p2 = p2
    if (87 >= -1) {
        p2 += p2
    }
    for fi4 := 0; fi4 < 8; fi4++ {
        p2 = (p2 << ((42) & 63))
        _ = fi4
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -141
    var v3 uint = 660
    v5 := GenS1{F0: false}
    if false {
        for fi6 := 0; fi6 < 5; fi6++ {
            for fi7 := 0; fi7 < 17; fi7++ {
                switch fi6 {
                case 0:
                    v2 = (v2 & (v2 % ((v2) | 1)))
                }
                var v8 uint16 = 42
                _ = len(v1)
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        if true {
            if true {
                var v9 uint8 = 42
                v10 := []int{-1, 1, 127, -166}
                _ = v10
                _ = v9
            }
            v2 = v2
        }
        v11 := GenS1{F0: (-2.25 <= (-2.25 + 1.5))}
        _ = v11
    }
    v12 := GenS2{F0: v2, F1: int(v2)}
    var v13 uint8 = 239
    v14 := []int{9223372036854775807, 9223372036854775807, 128, -1}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v15 := GenS2{F0: 42, F1: -9223372036854775808}
    for fi16 := 0; fi16 < 19; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            if true {
                var v18 uint32 = 2
                _ = len(v14)
                _ = v18
            } else {
                v19 := GenS1{F0: false}
                v2 += v2
                fmt.Println(fi17)
                fmt.Println(fi17, fi17, fi17)
                v20 := GenS3{F0: 9223372036854775807, F1: 509, F2: false}
                _ = v19
                _ = v20
            }
            var v21 int16 = 1000
            var v22 int64 = -1
            _ = fi17
            _ = v21
            _ = v22
        }
        switch fi16 {
        case 127:
            var v23 int64 = 1
            _ = v23
        }
        _ = fi16
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 uint = 65535
        for fi26 := 0; fi26 < 14; fi26++ {
            v27 := GenS2{F0: 81, F1: (fi26 >> ((255) & 63))}
            delete(v1, "a")
            v28 := GenS1{F0: (v13 != (v13 & v13))}
            _ = fi26
            _ = v27
            _ = v28
        }
        v29 := GenS4{F0: v13}
        _ = fi24
        _ = v25
        _ = v29
    }
    switch 127 {
    case 127:
        v1["hello"] = -9223372036854775808
    }
    if false {
        if true {
            if (true || (1e10 >= 3.14159)) {
                var v30 uint8 = (^(v13))
                if true {
                    switch 32768 {
                    case -958:
                        v31 := v1["hello"]
                        _ = v31
                    default:
                        if ((18446744073709551615 == v3) || (32768 != 0)) {
                            if false {
                                var v32 string = ("foo bar" + "a")
                                _ = v32
                            } else {
                                switch 2147483647 {
                                case 32768:
                                    var v33 uint8 = v30
                                    _ = v33
                                case 9223372036854775807:
                                    v30 = (v13 * v30)
                                case -228:
                                    v13 -= v13
                                }
                                var v34 int16 = 42
                                var v35 bool = ((v2 >= 42) || false)
                                _ = v34
                                _ = v35
                            }
                            v3 = (v3 >> ((65535) & 63))
                            v30 = v30
                        } else {
                            v36 := GenS5{F0: 0}
                            var v37 uint8 = (v13 << (((v3 ^ v3)) & 63))
                            v30 += v30
                            _ = v36
                            _ = v37
                        }
                    }
                    if (false || ((904 > 1) || false)) {
                        var v38 string = "hello"
                        var v39 uint8 = 154
                        _ = v38
                        _ = v39
                    }
                    v14 = append(v14, 42)
                }
                _ = v30
            } else {
                if true {
                    var v40 uint32 = 4294967295
                    var v41 uint = (v3 ^ (^(v3)))
                    _ = v40
                    _ = v41
                }
            }
        }
    }
    v42 := GenS1{F0: false}
    fmt.Println(v13)
    v2 -= -2147483648
    _ = len(v14)
    fmt.Println(v13, v13, v13)
    switch "" {
    case "a":
        v3 -= v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v42
    _ = v5
}
