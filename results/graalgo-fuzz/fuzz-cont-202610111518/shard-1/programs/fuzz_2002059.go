package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int32; F2 int64 }

func main() {
    var v1 int16 = 0
    var v2 int64 = 9223372036854775807
    v3 := []string{}
    var v4 int = 127
    var v5 bool = true
    var v6 uint8 = 42
    var v7 float32 = ((1.5 / 1.0) + 0.1)
    var v8 int32 = -382
    v8 = (v8 >> ((772) & 63))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: 2147483647, F1: (77 < v6)}
        _ = len(v3)
        _ = fi9
        _ = v10
    }
    if v5 {
        switch (^(v4)) {
        case -570:
            var v11 uint16 = 586
            _ = v11
        case 2147483647:
            for fi12 := 0; fi12 < 6; fi12++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi12
            }
        case -2147483648:
            switch v4 {
            case -958:
                v13 := GenS1{F0: 2147483647, F1: (1 != 0)}
                _ = v13
            case -548:
                var v14 uint32 = 515
                _ = v14
            case 4294967295:
                if (!(!v5)) {
                    fmt.Println(v1, v1, v7)
                    v2 = v2
                    switch "" {
                    case "a":
                        fmt.Println(v7, v1)
                    case "foo bar":
                        var v15 int8 = -54
                        _ = v15
                    default:
                        for fi16 := 0; fi16 < 3; fi16++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi16
                        }
                    }
                } else {
                    var v17 bool = v5
                    var v18 uint64 = 18446744073709551615
                    if (v5 && (!(((v6 * v6) != 1) || (((v4 >= int(v18)) && v5) && (v2 != (v2 + -249)))))) {
                        var v19 int32 = (v8 ^ (^(v8)))
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 17; fi20++ {
                            var v21 uint64 = v18
                            v6 += v6
                            _ = fi20
                            _ = v21
                        }
                        var v22 int = v4
                        var v23 int64 = (v2 ^ v2)
                        _ = v22
                        _ = v23
                    }
                    _ = v17
                    _ = v18
                }
            }
        }
    }
    var v24 bool = (255 != 891)
    fmt.Println(v2)
    fmt.Println(v8)
    v25 := GenS2{F0: 100}
    for fi26 := 0; fi26 < 9; fi26++ {
        if (((32767 > (-(v1))) && (!v5)) || true) {
            switch (fi26 - (fi26 - fi26)) {
            case 32768:
                _ = len(v3)
            case 128:
                var v27 uint64 = 827
                _ = v27
            default:
                v3 = append(v3, (("\n\t" + ("" + "hello")) + "hello"))
            }
            switch ("a" + "\n\t") {
            case "a":
                var v28 uint32 = 898
                _ = v28
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                switch -12 {
                case 128:
                    v4 = -1
                default:
                    var v29 float64 = 0.0
                    _ = v29
                }
            }
        }
        _ = fi26
    }
    var v30 int8 = 96
    v30 = v30
    v31 := GenS3{F0: "", F1: 139, F2: -1}
    for fi32 := 0; fi32 < 20; fi32++ {
        v33 := GenS2{F0: -39}
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
