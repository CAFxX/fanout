package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }

func main() {
    var v1 bool = ((!true) && (!(42 > 689)))
    v2 := []uint32{}
    var v3 uint8 = 42
    var v4 int = -2147483648
    var v5 uint = 18446744073709551615
    var v6 uint32 = 2147483648
    if (((!(2 <= uint64(v5))) && v1) || v1) {
        var v7 bool = (-579 == v4)
        if v7 {
            fmt.Println(v6, v1)
        } else {
            switch 0 {
            case 0:
                for fi8 := 0; fi8 < 18; fi8++ {
                    var v9 int32 = 1
                    switch "foo bar" {
                    case "hello":
                        var v10 uint16 = 42
                        _ = v10
                    case "a":
                        if (v5 < v5) {
                            var v11 uint16 = 533
                            var v12 int = fi8
                            _ = v11
                            _ = v12
                        }
                    case "\n\t":
                        v13 := GenS1{F0: 440, F1: -1}
                        _ = v13
                    }
                    _ = fi8
                    _ = v9
                }
            case 746:
                var v14 uint = 827
                _ = v14
            }
            var v15 int = v4
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int64 = -678
            var v18 bool = (float32(0.0) != 0.1)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v7
    }
    v5 -= (^(v5))
    for fi19 := 0; fi19 < 15; fi19++ {
        if (!false) {
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 uint32 = v6
                var v22 int8 = 91
                if (uint64(v3) >= 4294967295) {
                    v23 := GenS1{F0: (fi19 >> (((v5 << ((v5) & 63))) & 63)), F1: -632}
                    _ = v23
                }
                _ = fi20
                _ = v21
                _ = v22
            }
            var v24 uint64 = 171
            for fi25 := 0; fi25 < 4; fi25++ {
                v6 = (^(v6))
                _ = fi25
            }
            _ = v24
        } else {
            var v26 uint8 = v3
            _ = v26
        }
        _ = fi19
    }
    v6 = v6
    for fi27 := 0; fi27 < 18; fi27++ {
        fmt.Println(v1, fi27)
        _ = fi27
    }
    v28 := GenS1{F0: (v4 << ((65535) & 63)), F1: 42}
    v3 -= v3
    fmt.Println(v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
