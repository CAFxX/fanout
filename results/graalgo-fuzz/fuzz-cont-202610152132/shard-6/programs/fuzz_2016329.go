package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint8 }
func genFn0(p0 int32, p1 string) int16 {
    for fi5 := 0; fi5 < 1; fi5++ {
        p1 = p1
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    p1 = ("hello" + p1)
    if false {
        switch p1 {
        case "hello":
            var v7 uint8 = 2
            _ = v7
        case "":
            v8 := []int{42, 0, -593}
            _ = v8
        }
    }
    if true {
        if true {
            var v9 uint8 = 1
            _ = v9
        } else {
            v10 := []int{-9223372036854775808}
            p1 = "foo bar"
            _ = v10
        }
    } else {
        var v11 int64 = -1
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 uint8 = 127
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: fi12, F1: 0}
        _ = fi12
        _ = v15
    }
    _ = p0
    _ = p1
    return 1000
}

func main() {
    var v1 uint = 65535
    v2 := []string{""}
    var v3 bool = false
    v4 := make(map[string]string)
    if v3 {
        v1 = 65535
        var v16 float64 = 3.14159
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v16
    } else {
        var v17 int64 = 0
        var v18 uint64 = 0
        if true {
            v17 -= (v17 & (v17 / ((v17) | 1)))
            v19 := GenS1{F0: 2147483647, F1: 89}
            if false {
                v20 := GenS2{F0: (v18 & v18), F1: 2147483647, F2: 128}
                var v21 uint32 = 4294967295
                var v22 string = ("a" + "hello")
                _ = v20
                _ = v21
                _ = v22
            } else {
                v1 = v1
                var v23 int16 = 851
                v3 = v3
                _ = v23
            }
            _ = v19
        } else {
            if len(v2) > 0 {
                v2[0] = ("" + "a")
            }
            for fi24 := 0; fi24 < 10; fi24++ {
                v18 += 42
                _ = fi24
            }
        }
        _ = v17
        _ = v18
    }
    var v25 uint16 = 65535
    fmt.Println(v1, v25, v3)
    v26 := GenS1{F0: -1, F1: 4294967295}
    v3 = v3
    for fi27 := 0; fi27 < 2; fi27++ {
        v3 = (!(true && ((-606 > int64(fi27)) || v3)))
        fmt.Println(v1, v25, v25)
        _ = fi27
    }
    v28 := v4[((((("a" + "hello") + (("" + ("foo bar" + "")) + ("" + "\n\t"))) + (("\n\t" + ("" + "a")) + ("a" + "\n\t"))) + "a") + "foo bar")]
    v29 := GenS1{F0: -15, F1: 0}
    v28 = v28
    if ((-1 != 42) || v3) {
        if v3 {
            switch "" {
            case "foo bar":
                switch ("a" + (v28 + v28)) {
                case "a":
                    for fi30 := 0; fi30 < 20; fi30++ {
                        var v31 uint64 = 2
                        if len(v2) > 0 {
                            v2[0] = ("hello" + v28)
                        }
                        if (v25 >= (v25 % ((v25) | 1))) {
                            delete(v4, (v28 + (v28 + (((v28 + "a") + v28) + v28))))
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                        } else {
                            if true {
                                if v3 {
                                    v3 = false
                                } else {
                                    var v32 uint16 = v25
                                    switch (fi30 << ((v1) & 63)) {
                                    case -526:
                                        for fi33 := 0; fi33 < 5; fi33++ {
                                            var v34 uint8 = 164
                                            fmt.Println(fi30)
                                            _ = fi33
                                            _ = v34
                                        }
                                    case 0:
                                        var v35 uint64 = v31
                                        _ = v35
                                    case -752:
                                        v3 = (false || (v3 && ((0.1 != ((0.1 * 1e10) + 3.14159)) || v3)))
                                    default:
                                        var v36 uint16 = v25
                                        _ = v36
                                    }
                                    _ = v32
                                }
                                var v37 uint64 = uint64(v1)
                                v2 = append(v2, "hello")
                                _ = v37
                            } else {
                                v4[(v28 + "hello")] = v28
                            }
                        }
                        _ = fi30
                        _ = v31
                    }
                default:
                    var v38 uint32 = 0
                    _ = v38
                }
            default:
                switch 945 {
                case 9223372036854775807:
                    v39 := GenS1{F0: 2147483647, F1: uint32(v1)}
                    _ = v39
                case -2147483648:
                    v1 = v1
                case -293:
                    v3 = v3
                default:
                    v40 := GenS2{F0: 1, F1: 65535, F2: 128}
                    _ = v40
                }
            }
            var v41 int64 = -1
            _ = v41
        } else {
            var v42 float32 = (3.14159 + 1.5)
            _ = v42
        }
        v28 = (v28 + (v28 + ("\n\t" + v28)))
        var v43 int8 = 1
        _ = v43
    } else {
        var v44 uint8 = 0
        _ = v44
    }
    var v45 uint32 = 1
    v25 = uint16(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v45
}
