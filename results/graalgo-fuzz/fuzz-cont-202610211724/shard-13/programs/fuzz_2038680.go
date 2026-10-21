package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 string, p1 uint16) bool {
    v4 := GenS1{F0: (p1 & (p1 ^ (p1 | (p1 << ((772) & 63)))))}
    if (!(!(834 >= 255))) {
        p0 = p0
        var v5 int16 = 795
        _ = v5
    } else {
        p0 = p0
    }
    v6 := []int{9223372036854775807, -9223372036854775808, 32767, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return false
}
func genFn1(p0 bool) uint64 {
    var v7 int32 = -1
    v8 := []int{42, 771, -9223372036854775808}
    v7 = 14
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint32 = 4294967295
    var v2 float64 = 0.1
    v3 := []int{}
    v9 := GenS1{F0: 65535}
    var v10 bool = (2 <= 42)
    v11 := GenS1{F0: 255}
    for fi12 := 0; fi12 < 8; fi12++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi12
    }
    var v13 uint = 265
    v14 := GenS1{F0: 256}
    if true {
        var v15 string = (("foo bar" + "\n\t") + "foo bar")
        _ = v15
    }
    var v16 bool = false
    if ((("a" + "\n\t") + ((("a" + "a") + "") + "\n\t")) >= "") {
        var v17 uint32 = v1
        v17 += (v1 - v17)
        var v18 float32 = ((3.14159 - ((float32(-2.25) + (3.14159 / 3.14159)) / 1e10)) + 0.0)
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
