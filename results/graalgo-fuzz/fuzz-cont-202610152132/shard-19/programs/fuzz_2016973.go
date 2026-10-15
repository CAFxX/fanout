package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool; F1 string; F2 int32 }

func main() {
    var v1 int32 = -92
    var v2 int8 = 39
    var v3 uint16 = 339
    v4 := make(map[uint32]bool)
    var v5 uint32 = uint32(v3)
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint = 65535
        v9 := []int{fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    v1 = (v1 - (v1 & int32(v2)))
    v6 = v6
    if (true && (!true)) {
        v10 := v4[(v5 - v5)]
        _ = v10
    }
    switch "" {
    case "":
        var v11 int16 = 1000
        _ = v11
    case "\n\t":
        switch ((("hello" + ("" + (("\n\t" + (("" + ("" + ("hello" + "foo bar"))) + "")) + "\n\t"))) + ("foo bar" + "\n\t")) + (("hello" + ("\n\t" + "")) + "foo bar")) {
        case "foo bar":
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 int = fi12
                _ = fi12
                _ = v13
            }
        }
    }
    v14 := []int{}
    v6 += (v6 % ((v6) | 1))
    v15 := GenS1{F0: ("a" + "\n\t"), F1: int32(v6), F2: 1}
    v16 := GenS2{F0: false, F1: "\n\t", F2: (v1 - v1)}
    _ = len(v14)
    fmt.Println(v1, v5)
    v17 := GenS1{F0: ("a" + "a"), F1: (-(v1)), F2: 1}
    var v18 int = 0
    if true {
        delete(v4, v5)
    }
    switch v18 {
    case 169:
        if (!((v18 % ((v18) | 1)) == 0)) {
            var v19 uint8 = (^(v6))
            switch -763 {
            case -9223372036854775808:
                delete(v4, (v5 ^ 42))
            case 9223372036854775807:
                switch int(v5) {
                case 32767:
                    var v20 int16 = int16(v3)
                    _ = v20
                case 4294967295:
                    if false {
                        var v21 int8 = (^(v2))
                        for fi22 := 0; fi22 < 11; fi22++ {
                            v6 -= v6
                            _ = fi22
                        }
                        if len(v14) > 0 {
                            v14[0] = (v18 >> ((2) & 63))
                        }
                        _ = v21
                    } else {
                        v3 += 2
                        var v23 bool = (v3 != v3)
                        _ = v23
                    }
                }
            default:
                v1 = (v1 << ((uint(v5)) & 63))
            }
            fmt.Println(v1, v3)
            _ = v19
        } else {
            for fi24 := 0; fi24 < 17; fi24++ {
                for fi25 := 0; fi25 < 10; fi25++ {
                    fmt.Println(v2, v1)
                    if (true && (true || true)) {
                        fmt.Println(v5, fi25)
                    }
                    _ = fi25
                }
                _ = fi24
            }
        }
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
