package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 int16 = 32767
    var v4 int = 0
    var v5 bool = true
    v6 := GenS1{F0: (("foo bar" + ("hello" + (("" + "\n\t") + ("\n\t" + "hello")))) != "foo bar")}
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 4; fi8++ {
        v3 = (v3 + (v3 + 32767))
        _ = len(v1)
        if false {
            var v9 uint16 = 255
            v7 -= 1e10
            v4 -= fi8
            _ = v9
        } else {
            if v5 {
                var v10 int32 = int32(v3)
                fmt.Println(fi8, fi8, v3)
                v7 += v7
                _ = v10
            }
            var v11 int32 = -2147483648
            _ = v11
        }
        _ = fi8
    }
    var v12 float32 = (v7 + v7)
    v7 -= (v12 - 1.5)
    v1["\n\t"] = v5
    v13 := []int{}
    var v14 uint8 = 0
    var v15 int = v4
    v12 += (float32(((0.1 + float64(-2.25)) + 3.14159)) * v7)
    delete(v2, "foo bar")
    var v16 uint8 = (v14 >> ((4294967295) & 63))
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 uint = 255
        var v19 uint32 = 2
        v20 := GenS1{F0: v5}
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    v5 = (!true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
