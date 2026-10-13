package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64, p1 bool, p2 int32) int16 {
    for fi4 := 0; fi4 < 5; fi4++ {
        p1 = p1
        if true {
            for fi5 := 0; fi5 < 8; fi5++ {
                v6 := GenS1{F0: p1, F1: 2147483647, F2: 127}
                fmt.Println(p0, p2)
                v7 := []int{fi4, int(p2), (^(fi5)), -9223372036854775808}
                var v8 int = (-(fi4))
                fmt.Println(fi4, p2)
                _ = fi5
                _ = v6
                _ = v7
                _ = v8
            }
            v9 := GenS2{F0: -181}
            var v10 int32 = (p2 / (((p2 ^ 0)) | 1))
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := []int{-2147483648, (fi4 | int(p2)), (fi11 >> ((65535) & 63))}
            p0 = p0
            _ = len(v12)
            _ = fi11
            _ = v12
        }
        _ = fi4
    }
    var v13 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    var v1 string = ("foo bar" + ("a" + (("a" + ((("\n\t" + "a") + "a") + "foo bar")) + "\n\t")))
    var v2 bool = false
    var v3 uint8 = 51
    if true {
        fmt.Println(v3, v1, v3)
        v14 := []int{12, 0, 42}
        _ = v14
    }
    v15 := GenS2{F0: -32768}
    v16 := GenS2{F0: genFn0(1, v2, 0)}
    v17 := GenS3{F0: 196}
    var v18 bool = (32768 != 793)
    var v19 string = (v1 + v1)
    var v20 int8 = -29
    var v21 int64 = 9223372036854775807
    v22 := GenS2{F0: -630}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
