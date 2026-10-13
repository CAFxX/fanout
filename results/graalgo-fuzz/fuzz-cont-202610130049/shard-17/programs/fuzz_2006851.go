package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 bool, p1 uint16) int16 {
    fmt.Println(p1, p0)
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(p1)
        _ = fi6
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 uint8, p1 int8) int8 {
    var v7 uint32 = 2
    fmt.Println(p0)
    v8 := []int{9223372036854775807, -39}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int = 32767
    var v2 int16 = 32767
    var v3 uint = 18446744073709551615
    var v4 uint64 = 0
    var v5 int8 = 42
    var v9 uint = v3
    v2 = genFn0((v3 != (v9 >> ((v3) & 63))), uint16(v9))
    v5 = v5
    v10 := GenS1{F0: 197}
    for fi11 := 0; fi11 < 2; fi11++ {
        v4 -= (v4 ^ (v4 | 968))
        _ = fi11
    }
    v9 = (v9 & v9)
    switch -979 {
    case 32768:
        v9 -= 974
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        switch (("hello" + ("" + "")) + "a") {
        case "hello":
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := []int{fi13}
                _ = fi13
                _ = v14
            }
        case "a":
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 int32 = -2147483648
                v5 = genFn1(2, v5)
                _ = fi15
                _ = v16
            }
        }
        var v17 uint64 = v4
        _ = fi12
        _ = v17
    }
    v1 += v1
    v3 -= (v3 & (v3 & 42))
    v18 := []int{(v1 ^ 32768), v1, v1}
    var v19 int64 = 9223372036854775807
    v20 := GenS1{F0: -967}
    if true {
        var v21 uint64 = v4
        for fi22 := 0; fi22 < 17; fi22++ {
            if true {
                v9 += 1
                var v23 float64 = ((0.1 / 0.1) / -2.25)
                for fi24 := 0; fi24 < 13; fi24++ {
                    if false {
                        fmt.Println(v4, v1)
                        if len(v18) > 0 {
                            v18[0] = -881
                        }
                        var v25 uint = v3
                        _ = v25
                    }
                    v26 := GenS1{F0: (v19 ^ (^(v19)))}
                    var v27 uint16 = 0
                    _ = fi24
                    _ = v26
                    _ = v27
                }
                _ = v23
            }
            _ = fi22
        }
        v1 -= int(v19)
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
