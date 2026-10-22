package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 int32 }
type GenS4 struct { F0 string; F1 int64 }

func main() {
    var v1 int16 = 1000
    v2 := []int{127, 9223372036854775807, 216, 148}
    var v3 uint64 = 42
    v1 = v1
    switch 32767 {
    case 42:
        v4 := GenS1{F0: 127, F1: 42, F2: v3}
        _ = v4
    }
    var v5 int16 = v1
    if (3.14159 != 1e10) {
        v6 := GenS2{F0: 0}
        var v7 int = 32767
        v8 := GenS3{F0: 870, F1: (v7 >> ((4294967295) & 63)), F2: 350}
        _ = v6
        _ = v7
        _ = v8
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v9 uint32 = 2147483648
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS1{F0: 100, F1: 206, F2: (v3 % ((42) | 1))}
        if true {
            v12 := GenS3{F0: 0, F1: (fi10 | fi10), F2: 2147483647}
            for fi13 := 0; fi13 < 10; fi13++ {
                v1 -= v1
                _ = fi13
            }
            _ = v12
        } else {
            v5 = v5
        }
        _ = fi10
        _ = v11
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v14 float64 = ((((float64((0.0 - 1e10)) * float64(((3.14159 + (0.0 / (0.1 + 0.0))) - 1.5))) + float64(1e10)) + 3.14159) * 1.5)
    switch -393 {
    case -555:
        if (32768 >= 2147483647) {
            var v15 int8 = 1
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 float64 = v14
                v18 := GenS4{F0: (("\n\t" + ("\n\t" + ("\n\t" + "foo bar"))) + "foo bar"), F1: 628}
                var v19 int = (fi16 | 2147483647)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = v15
        } else {
            if (true || true) {
                v20 := GenS1{F0: -47, F1: uint16(v9), F2: (^(v3))}
                v3 = uint64(v9)
                fmt.Println(v9, v14, v14)
                _ = v20
            }
            var v21 uint32 = (v9 * v9)
            _ = v21
        }
    }
    v3 = (v3 + v3)
    if false {
        if (uint16(v3) < 1) {
            _ = len(v2)
            v2 = append(v2, 9223372036854775807)
            if false {
                var v22 int8 = 100
                _ = v22
            } else {
                if len(v2) > 0 {
                    v2[0] = 420
                }
            }
        }
    }
    var v23 int16 = (v5 * (v1 << ((uint(v3)) & 63)))
    switch 2147483647 {
    case 32768:
        for fi24 := 0; fi24 < 2; fi24++ {
            switch (-(fi24)) {
            case 128:
                v1 += (v1 / ((-824) | 1))
            case -9223372036854775808:
                var v25 uint64 = 18446744073709551615
                _ = v25
            }
            _ = fi24
        }
    case -303:
        fmt.Println(v9)
    }
    var v26 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v5
    _ = v9
}
