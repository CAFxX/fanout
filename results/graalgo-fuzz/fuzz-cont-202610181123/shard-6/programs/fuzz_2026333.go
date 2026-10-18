package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 string; F1 uint32; F2 int }
func genFn0(p0 uint) int {
    for fi7 := 0; fi7 < 9; fi7++ {
        switch ("foo bar" + "hello") {
        case "":
            if true {
                v8 := []int{fi7, 860, fi7, fi7}
                if len(v8) > 0 {
                    v8[0] = (fi7 << ((784) & 63))
                }
                _ = v8
            }
        case "a":
            v9 := []int{fi7, (fi7 - fi7), (fi7 - -2147483648)}
            _ = v9
        case "foo bar":
            var v10 bool = ("\n\t" <= "")
            _ = v10
        }
        var v11 string = (("foo bar" + ("hello" + "hello")) + "hello")
        _ = fi7
        _ = v11
    }
    var v12 int8 = 100
    var v13 uint32 = 2147483647
    _ = p0
    _ = v12
    _ = v13
    return 290
}

func main() {
    var v1 uint64 = 2
    var v2 uint64 = v1
    var v3 uint16 = 255
    var v4 uint64 = (^(v2))
    var v5 uint32 = 42
    var v6 uint64 = (v2 << ((2) & 63))
    v3 -= (v3 - (v3 | 183))
    v14 := GenS1{F0: 2, F1: 244}
    v15 := GenS2{F0: "a", F1: v5, F2: 127}
    v16 := []int{933, -9223372036854775808, genFn0(226)}
    var v17 int32 = 990
    fmt.Println(v1, v2)
    for fi18 := 0; fi18 < 13; fi18++ {
        v17 += v17
        v19 := GenS1{F0: v6, F1: 183}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        if false {
            v6 += v1
            v2 -= 18446744073709551615
        }
        _ = fi20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v17 = 1
        var v22 int32 = (v17 / (((v17 - -2147483648)) | 1))
        v5 += 671
        _ = fi21
        _ = v22
    }
    fmt.Println(v17, v5, v17)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
