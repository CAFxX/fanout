package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int8; F2 int8 }
func genFn0(p0 int64, p1 uint64, p2 uint32) int32 {
    fmt.Println(p0, p0)
    v4 := GenS1{F0: 32767, F1: p2, F2: -647}
    for fi5 := 0; fi5 < 14; fi5++ {
        if true {
            v6 := GenS1{F0: 0, F1: 822, F2: -2147483648}
            _ = v6
        }
        _ = fi5
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}

func main() {
    var v1 uint = 2
    v2 := []int{128, 9223372036854775807}
    var v3 float32 = (0.0 * 3.14159)
    var v7 int64 = 0
    var v8 uint = 2
    v9 := GenS1{F0: 761, F1: 2147483648, F2: 555}
    if true {
        v3 = v3
    }
    v7 = v7
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS2{F0: 1, F1: 127}
            var v13 bool = (1 <= 42)
            v13 = (!v13)
            v14 := GenS2{F0: 4294967295, F1: 227}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        if true {
            if true {
                for fi15 := 0; fi15 < 5; fi15++ {
                    v1 += 65535
                    var v16 int = fi10
                    _ = fi15
                    _ = v16
                }
            }
            switch "\n\t" {
            case "hello":
                fmt.Println(fi10)
            case "foo bar":
                var v17 uint = 16
                _ = v17
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                v3 -= ((1.5 / (((v3 + 1.5) / v3) - v3)) + v3)
                fmt.Println(v8, fi18)
                _ = fi18
            }
        }
        if false {
            if (2147483647 <= genFn0(v7, 2, 2)) {
                fmt.Println(v1, v1, v1)
            } else {
                if len(v2) > 0 {
                    v2[0] = fi10
                }
                fmt.Println(v7)
            }
            var v19 int8 = 1
            switch fi10 {
            case 42:
                if (float64(-2.25) < -2.25) {
                    var v20 int64 = 409
                    if len(v2) > 0 {
                        v2[0] = fi10
                    }
                    var v21 uint8 = 2
                    _ = v20
                    _ = v21
                }
            }
            _ = v19
        }
        _ = fi10
    }
    var v22 string = ""
    v23 := GenS1{F0: 1, F1: 948, F2: 939}
    fmt.Println(v22, v8)
    v24 := GenS3{F0: 250, F1: 1, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
