package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint8 = 92
    v3 := []uint32{403, (^(v1)), (^(v1)), (^(v1))}
    v4 := make(map[int]string)
    v5 := make(map[uint32]bool)
    var v6 uint = 425
    v5[v1] = true
    v1 = v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, (^(v1)))
    if true {
        for fi7 := 0; fi7 < 5; fi7++ {
            v1 += v1
            if (862 < 42) {
                v8 := GenS1{F0: "hello"}
                _ = v8
            }
            v9 := GenS1{F0: "a"}
            _ = fi7
            _ = v9
        }
    }
    v10 := GenS2{F0: 0}
    v6 += v6
    if (((true || true) || false) && (42 != 540)) {
        switch -9223372036854775808 {
        case 2147483647:
            v1 += 2
        case 9223372036854775807:
            if (true || true) {
                v6 += (v6 & (^(v6)))
            }
        }
        for fi11 := 0; fi11 < 1; fi11++ {
            v5[v1] = true
            fmt.Println(v2, v6)
            _ = fi11
        }
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        if (!true) {
            v13 := GenS2{F0: 32767}
            for fi14 := 0; fi14 < 14; fi14++ {
                v1 -= v1
                v2 -= v2
                _ = fi14
            }
            _ = v13
        } else {
            switch "a" {
            case "a":
                if (!(!true)) {
                    v15 := GenS2{F0: -722}
                    _ = v15
                }
            }
        }
        switch "hello" {
        case "":
            for fi16 := 0; fi16 < 17; fi16++ {
                v6 -= 42
                var v17 uint32 = (v1 + (^(v1)))
                var v18 uint32 = (^(v17))
                _ = fi16
                _ = v17
                _ = v18
            }
        case "foo bar":
            _ = len(v3)
        case "a":
            var v19 uint8 = v2
            _ = v19
        default:
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 int = (fi20 % (((fi12 >> (((^(v6))) & 63))) | 1))
                var v22 int8 = 0
                v22 -= 54
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi12
    }
    v2 -= (v2 | v2)
    v3 = append(v3, (^(v1)))
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
