package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint }
type GenS3 struct { F0 uint; F1 uint; F2 uint }

func main() {
    var v1 uint64 = 740
    var v2 uint64 = (v1 | 4294967295)
    v3 := []int{0, -9223372036854775808, -495, -9223372036854775808}
    v4 := make(map[string]uint32)
    var v5 int32 = -2147483648
    var v6 uint16 = 0
    v5 -= 42
    var v7 int = -901
    var v8 int8 = 42
    var v9 bool = (("a" + ("" + "a")) != "foo bar")
    if v9 {
        if (9223372036854775807 != 52) {
            for fi10 := 0; fi10 < 19; fi10++ {
                switch (fi10 / (((fi10 - (-(fi10)))) | 1)) {
                case 9223372036854775807:
                    v5 += v5
                case 32767:
                    var v11 int = (^(fi10))
                    _ = v11
                case 32768:
                    v3 = append(v3, (-(fi10)))
                }
                switch ("hello" + (("a" + "hello") + "\n\t")) {
                case "hello":
                    if false {
                        fmt.Println(v6)
                    } else {
                        for fi12 := 0; fi12 < 19; fi12++ {
                            v1 += 4294967295
                            _ = fi12
                        }
                    }
                case "":
                    if (-804 != 0) {
                        switch v7 {
                        case -734:
                            v13 := GenS1{F0: (^(v6)), F1: 255}
                            _ = v13
                        case 127:
                            var v14 int64 = 671
                            _ = v14
                        case 48:
                            switch -1000 {
                            case 0:
                                v15 := GenS2{F0: 1000, F1: v6, F2: 4294967295}
                                _ = v15
                            case 32768:
                                if (((-2.25 - 1.5) + 3.14159) == 3.14159) {
                                    switch v7 {
                                    case 128:
                                        var v16 int = 32767
                                        _ = v16
                                    default:
                                        fmt.Println(v6, v7, v6)
                                    }
                                }
                            case 128:
                                fmt.Println(v8)
                            default:
                                v3 = append(v3, v7)
                            }
                        default:
                            v17 := GenS3{F0: 255, F1: 461, F2: 2}
                            _ = v17
                        }
                        switch "\n\t" {
                        case "foo bar":
                            if ((("\n\t" <= "hello") && (!(9223372036854775807 > -9223372036854775808))) && false) {
                                if v9 {
                                    var v18 uint32 = 1
                                    delete(v4, ("" + ((("hello" + ("" + (("foo bar" + ("hello" + ("hello" + ("a" + "a")))) + "hello"))) + "a") + ("foo bar" + ""))))
                                    _ = v18
                                } else {
                                    v8 -= 127
                                    switch "" {
                                    case "foo bar":
                                        v8 = v8
                                    case "hello":
                                        fmt.Println(v6, v2, v9)
                                    }
                                }
                                var v19 int = v7
                                _ = v19
                            }
                        case "hello":
                            if (1 <= 42) {
                                v1 -= 2
                            }
                        case "":
                            var v20 uint16 = 65535
                            _ = v20
                        }
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = v7
                    }
                }
                _ = fi10
            }
            var v21 bool = (9223372036854775807 < -927)
            fmt.Println(v2)
            _ = v21
        } else {
            v6 += (v6 | v6)
            v3 = append(v3, (v7 << ((348) & 63)))
            v9 = ((3.14159 - 1e10) == ((1e10 + 0.1) + (0.1 * (float64(0.0) / 1e10))))
        }
        v1 = v1
    } else {
        v7 += (v7 * (v7 * (^(v7))))
    }
    var v22 int32 = (v5 >> ((1) & 63))
    v23 := GenS1{F0: 42, F1: 0}
    var v24 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
