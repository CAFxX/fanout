package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint32, p1 uint8) uint {
    p0 -= (p0 >> ((65535) & 63))
    fmt.Println(p1)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 string = (("foo bar" + ("a" + "\n\t")) + (("foo bar" + "a") + "foo bar"))
    v2 := make(map[string]uint32)
    v3 := []string{v1, v1}
    var v4 int64 = 434
    var v5 int = 32768
    var v6 int16 = 32767
    for fi7 := 0; fi7 < 7; fi7++ {
        if ((!(!false)) && false) {
            v4 += v4
        }
        if (!(v4 < v4)) {
            fmt.Println(v1, v6, v6)
            v8 := GenS1{F0: 1, F1: 767, F2: false}
            fmt.Println(v6, v5, fi7)
            _ = v8
        }
        _ = fi7
    }
    v9 := GenS1{F0: 2, F1: 51, F2: false}
    fmt.Println(v6)
    switch (v5 - int(v4)) {
    case 32767:
        var v10 int16 = v6
        _ = v10
    }
    v11 := GenS1{F0: 2, F1: 337, F2: (42 > -1)}
    var v12 int16 = 32767
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS1{F0: 127, F1: 0, F2: ("foo bar" != v1)}
        v15 := GenS2{F0: 18446744073709551615}
        v5 -= fi13
        _ = fi13
        _ = v14
        _ = v15
    }
    v4 += -1
    fmt.Println(v4, v5)
    delete(v2, "a")
    v5 += v5
    v16 := GenS3{F0: 18446744073709551615, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
