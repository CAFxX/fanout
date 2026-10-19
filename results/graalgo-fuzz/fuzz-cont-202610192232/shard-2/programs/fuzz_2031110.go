package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint32) uint {
    v7 := GenS1{F0: -1}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v7
    return 49
}

func main() {
    var v1 float32 = 0.1
    v2 := []int{-311, 875, 128, 688}
    var v3 int = 1
    v4 := []uint32{1, 828}
    var v5 float32 = (float32(1e10) * v1)
    v6 := make(map[string]uint32)
    var v8 float32 = v5
    for fi9 := 0; fi9 < 19; fi9++ {
        v8 = v5
        _ = fi9
    }
    v10 := v6[("foo bar" + "foo bar")]
    v6["hello"] = v10
    v11 := GenS1{F0: v3}
    v10 += v10
    v1 = v1
    v3 = (v3 * v3)
    v5 -= ((v8 - v1) + v1)
    v5 = v1
    var v12 uint = genFn0(2, v10)
    v1 -= v5
    for fi13 := 0; fi13 < 9; fi13++ {
        if len(v4) > 0 {
            v4[0] = 735
        }
        if true {
            var v14 uint = v12
            _ = v14
        } else {
            v15 := GenS2{F0: (v3 << ((genFn0(32767, v10)) & 63))}
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 int8 = -1
                _ = fi16
                _ = v17
            }
            var v18 uint64 = 4294967295
            _ = v15
            _ = v18
        }
        _ = fi13
    }
    v6["hello"] = (^(v10))
    v19 := GenS3{F0: 0}
    for fi20 := 0; fi20 < 3; fi20++ {
        v12 = (v12 >> ((genFn0(255, (v10 / ((v10) | 1)))) & 63))
        _ = fi20
    }
    switch ("a" + "") {
    case "foo bar":
        v21 := GenS1{F0: (v3 - (v3 | (v3 << (((v12 ^ v12)) & 63))))}
        _ = v21
    case "\n\t":
        switch ("\n\t" + "foo bar") {
        case "hello":
            if len(v2) > 0 {
                v2[0] = 408
            }
        case "foo bar":
            fmt.Println(v10, v12)
        }
    case "":
        for fi22 := 0; fi22 < 7; fi22++ {
            if (!("a" >= "\n\t")) {
                if (0.0 >= v5) {
                    delete(v6, ("hello" + "hello"))
                    for fi23 := 0; fi23 < 16; fi23++ {
                        v24 := GenS2{F0: (v3 ^ v3)}
                        _ = fi23
                        _ = v24
                    }
                    v25 := v6["foo bar"]
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 18; fi26++ {
                v3 = v3
                _ = fi26
            }
            switch (fi22 * (fi22 ^ 4294967295)) {
            case 42:
                v10 = (v10 + 2147483648)
            case 0:
                for fi27 := 0; fi27 < 16; fi27++ {
                    var v28 uint16 = 720
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v12 = v12
                    _ = fi27
                    _ = v28
                }
            }
            _ = fi22
        }
    default:
        var v29 string = ""
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
