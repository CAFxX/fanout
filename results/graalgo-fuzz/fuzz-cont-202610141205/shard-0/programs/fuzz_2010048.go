package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 string; F2 int8 }

func main() {
    v1 := []string{"hello", "\n\t", "hello", "hello"}
    var v2 int = -953
    v3 := []string{"hello", "foo bar"}
    fmt.Println(v2)
    var v4 uint32 = 143
    v4 = (v4 + (v4 << ((42) & 63)))
    var v5 int32 = -626
    var v6 int64 = 536
    switch (("hello" + (("hello" + "a") + "a")) + ("hello" + "\n\t")) {
    case "a":
        v7 := GenS1{F0: v5, F1: "", F2: 18446744073709551615}
        _ = v7
    case "foo bar":
        if false {
            v4 = (v4 - (v4 - 0))
            v5 -= v5
        } else {
            var v8 uint32 = 1
            v3 = append(v3, "a")
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 string = "\n\t"
                var v11 float64 = (0.0 / 1.0)
                if true {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v4 = 2147483647
                        v10 = "\n\t"
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi12
                    }
                    _ = len(v1)
                } else {
                    fmt.Println(v2)
                    v13 := GenS2{F0: 42}
                    _ = v13
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v8
        }
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        if (!true) {
            v15 := GenS1{F0: int32(fi14), F1: "foo bar", F2: 0}
            var v16 uint64 = 885
            var v17 int32 = (v5 >> ((701) & 63))
            _ = v15
            _ = v16
            _ = v17
        } else {
            var v18 int = 2147483647
            for fi19 := 0; fi19 < 19; fi19++ {
                v6 = v6
                if false {
                    switch 1 {
                    case 4294967295:
                        v20 := GenS2{F0: v6}
                        _ = v20
                    case -2147483648:
                        var v21 uint16 = 42
                        _ = v21
                    case -71:
                        fmt.Println(fi19, v4)
                    default:
                        fmt.Println(v2, fi19, fi19)
                    }
                    if false {
                        v4 = v4
                        switch (fi14 | v2) {
                        case 128:
                            v22 := GenS3{F0: 65535}
                            _ = v22
                        default:
                            var v23 bool = false
                            _ = v23
                        }
                        if (81 == -67) {
                            v18 = 4294967295
                            switch v18 {
                            case -2147483648:
                                var v24 int16 = 1
                                _ = v24
                            default:
                                _ = len(v1)
                            }
                            switch v18 {
                            case 9223372036854775807:
                                var v25 int8 = 127
                                _ = v25
                            case 42:
                                v18 = fi14
                            case -492:
                                var v26 uint64 = 276
                                _ = v26
                            default:
                                v1 = append(v1, "\n\t")
                            }
                        } else {
                            v2 = fi14
                            v5 = int32(v6)
                            v3 = append(v3, (((("\n\t" + "a") + "a") + "foo bar") + ""))
                        }
                    }
                } else {
                    v27 := GenS1{F0: (v5 << ((18446744073709551615) & 63)), F1: ("\n\t" + ""), F2: 1}
                    v6 = v6
                    v5 = v5
                    _ = v27
                }
                _ = fi19
            }
            _ = v18
        }
        var v28 float32 = (0.1 + ((-2.25 * 3.14159) / ((0.1 - 0.0) - (float32((1.5 + (((0.1 - -2.25) * 0.1) - 3.14159))) - float32(((1.5 / (3.14159 - float64(0.1))) - ((0.0 / -2.25) - float64(float32((-2.25 - float64(((0.0 / (0.1 * 3.14159)) + 3.14159))))))))))))
        v28 = v28
        _ = fi14
        _ = v28
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v29 int64 = -428
    v29 -= -9223372036854775808
    fmt.Println(v4)
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 bool = false
        var v32 int16 = -1
        _ = fi30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        v34 := GenS4{F0: v6, F1: ("foo bar" + "\n\t"), F2: 42}
        _ = fi33
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
