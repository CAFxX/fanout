package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 48
    var v3 int32 = 42
    v4 := []string{}
    for fi5 := 0; fi5 < 3; fi5++ {
        switch (fi5 - 2147483647) {
        case 128:
            v6 := GenS1{F0: 65535}
            _ = v6
        default:
            if (635 != 255) {
                switch (fi5 >> ((255) & 63)) {
                case -819:
                    fmt.Println(fi5, v2, v3)
                case 9223372036854775807:
                    if false {
                        fmt.Println(v3)
                        var v7 bool = true
                        _ = v7
                    }
                }
            } else {
                v4 = append(v4, "")
                for fi8 := 0; fi8 < 5; fi8++ {
                    v9 := GenS1{F0: 65535}
                    _ = fi8
                    _ = v9
                }
                if false {
                    for fi10 := 0; fi10 < 2; fi10++ {
                        delete(v1, (fi10 * (fi5 & fi5)))
                        _ = fi10
                    }
                    var v11 float32 = (float32((float64(0.1) + (1.5 * float64(0.0)))) * (1.5 * ((-2.25 + -2.25) / (3.14159 + 1.5))))
                    var v12 int32 = v3
                    _ = v11
                    _ = v12
                }
            }
        }
        var v13 int64 = -9223372036854775808
        _ = fi5
        _ = v13
    }
    v2 -= 65535
    v14 := GenS1{F0: 42}
    v3 += 982
    switch (("hello" + "") + "foo bar") {
    case "foo bar":
        v2 = 65535
    default:
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := GenS1{F0: 0}
            delete(v1, fi15)
            _ = fi15
            _ = v16
        }
    }
    if len(v4) > 0 {
        v4[0] = ("a" + ("a" + ""))
    }
    var v17 float32 = float32(3.14159)
    var v18 int64 = 9223372036854775807
    var v19 uint = 4294967295
    var v20 uint32 = uint32(v19)
    v17 = 0.0
    for fi21 := 0; fi21 < 19; fi21++ {
        v3 = (v3 * (v3 ^ (v3 + v3)))
        var v22 int32 = 463
        v23 := GenS1{F0: 4294967295}
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
