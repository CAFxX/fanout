package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 uint64) uint32 {
    p0 = (^(p0))
    switch -896 {
    case 32768:
        var v5 int = 128
        _ = v5
    case 550:
        if true {
            v6 := []int{520, 32768}
            p0 -= (p0 & p0)
            _ = v6
        }
    }
    _ = p0
    return 274
}

func main() {
    var v1 uint = 2
    var v2 float64 = 1.5
    var v3 int = 938
    var v4 uint = (^(v1))
    var v7 int8 = 0
    var v8 uint8 = 1
    v4 += (v1 | (v1 - (v1 & 669)))
    v1 = 255
    fmt.Println(v4)
    for fi9 := 0; fi9 < 9; fi9++ {
        switch (("" + "a") + "") {
        case "":
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := []int{(-(fi9)), (fi10 >> ((v4) & 63))}
                v3 -= (fi9 - (fi10 >> (((^(v1))) & 63)))
                _ = fi10
                _ = v11
            }
        }
        v7 += (-(v7))
        _ = fi9
    }
    v12 := GenS1{F0: 18446744073709551615, F1: 4294967295}
    v13 := GenS2{F0: (("hello" + ((("a" + "") + "a") + "foo bar")) + "\n\t"), F1: 844}
    var v14 uint64 = 2
    if true {
        switch (v3 % ((4294967295) | 1)) {
        case -191:
            var v15 uint64 = (v14 % (((^(v14))) | 1))
            _ = v15
        case 242:
            v16 := GenS2{F0: ("foo bar" + "foo bar"), F1: 1}
            _ = v16
        case 2147483647:
            fmt.Println(v7, v4, v3)
        }
        fmt.Println(v14)
    }
    v2 = (0.1 - v2)
    if true {
        v17 := []int{(^(v3)), -2147483648, int(v1), (v3 & v3)}
        for fi18 := 0; fi18 < 20; fi18++ {
            if len(v17) > 0 {
                v17[0] = v3
            }
            _ = fi18
        }
        _ = v17
    }
    v4 = (v1 + v1)
    var v19 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
