package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("" + (("\n\t" + "a") + ("foo bar" + ("foo bar" + "foo bar"))))}
    var v3 string = "a"
    var v4 int = 0
    if false {
        for fi5 := 0; fi5 < 18; fi5++ {
            if (-232 >= 848) {
                fmt.Println(v4, v4, v4)
            } else {
                var v6 uint8 = 13
                _ = v6
            }
            v3 = "\n\t"
            _ = fi5
        }
        if (42 == 128) {
            v4 = v4
            v7 := v1[1]
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 bool = true
        _ = fi8
        _ = v9
    }
    delete(v1, v4)
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(fi10, v4, fi10)
        _ = fi10
    }
    var v11 int16 = -32768
    var v12 float64 = (((3.14159 + 0.0) - 0.1) + (float64(1.5) + (1.5 - 0.0)))
    v12 -= v12
    if true {
        delete(v1, -9223372036854775808)
    } else {
        for fi13 := 0; fi13 < 17; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 int16 = -737
                var v16 string = ("a" + (v3 + "\n\t"))
                v15 = v11
                if (v15 == (v15 % ((-316) | 1))) {
                    var v17 int8 = 46
                    _ = v17
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            if true {
                v11 -= (v11 / ((42) | 1))
                if (false && (!(42 >= 957))) {
                    v18 := GenS1{F0: 42}
                    fmt.Println(v4, v12)
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        v11 -= (v11 >> ((1) & 63))
                        _ = fi19
                    }
                }
            }
            _ = fi13
        }
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        if (true || false) {
            v4 -= fi20
        }
        var v21 int16 = -32768
        var v22 uint32 = 1
        _ = fi20
        _ = v21
        _ = v22
    }
    v11 = -437
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 string = v3
        _ = fi23
        _ = v24
    }
    v11 = -32768
    var v25 int64 = -468
    v1[(^(v4))] = 65535
    fmt.Println(v25, v3, v25)
    if (754 != 297) {
        v3 = ""
        var v26 int16 = v11
        v27 := GenS1{F0: 1}
        _ = v26
        _ = v27
    }
    v3 = v3
    var v28 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
