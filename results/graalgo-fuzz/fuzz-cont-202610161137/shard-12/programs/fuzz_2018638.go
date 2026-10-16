package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int }
func genFn0(p0 bool) uint32 {
    switch "" {
    case "foo bar":
        var v5 float64 = (0.1 * (1e10 * (1.5 / -2.25)))
        _ = v5
    case "hello":
        for fi6 := 0; fi6 < 12; fi6++ {
            if (56 != 65535) {
                var v7 int16 = 1000
                var v8 int32 = 2147483647
                if p0 {
                    v8 = (^(v8))
                    if (false || p0) {
                        fmt.Println(v8, fi6)
                        var v9 int64 = -1
                        v8 = v8
                        _ = v9
                    }
                } else {
                    var v10 bool = p0
                    var v11 int64 = int64(v7)
                    for fi12 := 0; fi12 < 14; fi12++ {
                        fmt.Println(v10, v8, v11)
                        switch -2147483648 {
                        case 128:
                            var v13 string = "a"
                            _ = v13
                        case 9223372036854775807:
                            v8 = v8
                        }
                        v14 := GenS1{F0: 2147483647, F1: fi6, F2: 0}
                        _ = fi12
                        _ = v14
                    }
                    _ = v10
                    _ = v11
                }
                _ = v7
                _ = v8
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                p0 = p0
                var v16 uint8 = 42
                v17 := GenS1{F0: 1, F1: fi6, F2: 808}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi6
        }
    default:
        switch "" {
        case "\n\t":
            v18 := GenS1{F0: 633, F1: 127, F2: 4294967295}
            _ = v18
        default:
            switch 32767 {
            case -374:
                v19 := GenS1{F0: 767, F1: 32768, F2: 167}
                _ = v19
            case -2147483648:
                if (false && true) {
                    var v20 uint64 = 616
                    _ = v20
                } else {
                    var v21 bool = (0 == 1)
                    _ = v21
                }
            case 2147483647:
                var v22 string = ((((("\n\t" + "a") + (("hello" + "\n\t") + "")) + "a") + ("a" + "foo bar")) + "")
                _ = v22
            }
        }
    }
    var v23 float64 = 0.1
    _ = p0
    _ = v23
    return 2
}

func main() {
    var v1 uint8 = 64
    v2 := make(map[int]int)
    var v3 int = 4294967295
    var v4 uint = 4294967295
    fmt.Println(v3, v4, v3)
    if false {
        v1 += (^(v1))
        fmt.Println(v4, v4)
        if true {
            v24 := GenS1{F0: 2147483647, F1: v3, F2: 1}
            v4 = v4
            _ = v24
        } else {
            switch v3 {
            case 0:
                for fi25 := 0; fi25 < 3; fi25++ {
                    v4 = 2
                    _ = fi25
                }
            case 32767:
                if false {
                    if (!(!((false && true) && false))) {
                        var v26 float32 = 1e10
                        fmt.Println(v4)
                        _ = v26
                    } else {
                        v3 -= 128
                    }
                } else {
                    v27 := v2[v3]
                    _ = v27
                }
            case 32768:
                fmt.Println(v1, v3)
            default:
                var v28 uint16 = 423
                _ = v28
            }
        }
    } else {
        _ = len(v2)
    }
    fmt.Println(v3)
    if (false && (true && false)) {
        if false {
            for fi29 := 0; fi29 < 18; fi29++ {
                fmt.Println(v4)
                _ = fi29
            }
        }
        if (65535 >= 42) {
            v1 = v1
            if (73 >= 730) {
                v30 := v2[(v3 + 871)]
                _ = v30
            }
        } else {
            v3 = v3
            for fi31 := 0; fi31 < 14; fi31++ {
                var v32 int32 = 89
                _ = fi31
                _ = v32
            }
        }
        if true {
            var v33 bool = true
            switch 149 {
            case -9223372036854775808:
                fmt.Println(v4)
            case -521:
                var v34 int = v3
                _ = v34
            case -1:
                switch 1 {
                case 127:
                    var v35 uint64 = 18446744073709551615
                    _ = v35
                case -1:
                    v2[v3] = (v3 % ((42) | 1))
                }
            default:
                v36 := GenS1{F0: 510, F1: 1, F2: 1}
                _ = v36
            }
            _ = v33
        }
    }
    var v37 uint64 = 94
    if false {
        var v38 int = v3
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        var v40 int8 = 127
        _ = fi39
        _ = v40
    }
    v41 := GenS2{F0: genFn0((false && false)), F1: (v3 - (v3 ^ v3))}
    v4 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
}
