package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint32, p1 int32) int32 {
    if (155 == 101) {
        switch 323 {
        case 0:
            p0 = p0
        case 9223372036854775807:
            v4 := GenS1{F0: 1}
            _ = v4
        }
        switch -1 {
        case 32767:
            fmt.Println(p1)
        }
    } else {
        p0 = p0
    }
    var v5 uint = 63
    _ = p0
    _ = p1
    _ = v5
    return p1
}

func main() {
    var v1 string = ("\n\t" + ("\n\t" + ""))
    v2 := []uint32{283}
    v3 := []string{(v1 + "foo bar"), v1, v1, v1}
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 255}
        _ = fi6
        _ = v7
    }
    var v8 int8 = -1
    if len(v2) > 0 {
        v2[0] = 1
    }
    v9 := GenS1{F0: 11}
    var v10 int16 = -669
    if false {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v10 += 32767
    } else {
        for fi11 := 0; fi11 < 3; fi11++ {
            if len(v3) > 0 {
                v3[0] = v1
            }
            v10 -= v10
            v1 = v1
            _ = fi11
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS1{F0: 161}
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS2{F0: (-9223372036854775808 < -1)}
    if true {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 int8 = v8
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            v8 = (v8 | (v8 << ((0) & 63)))
            _ = fi17
        }
    }
    var v18 int32 = -2147483648
    var v19 uint16 = 738
    v20 := GenS2{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
