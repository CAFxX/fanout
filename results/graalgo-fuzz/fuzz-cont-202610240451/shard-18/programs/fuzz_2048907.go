package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int = 32767
    v2 := make(map[int]bool)
    v3 := []int{(v1 & 42), (v1 & (^(v1))), (v1 * v1), 128}
    var v4 int8 = 42
    v5 := []string{"\n\t", "hello"}
    fmt.Println(v1, v1)
    v1 -= v1
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 string = ""
        v8 := GenS1{F0: 42}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v1, v1)
        _ = fi9
    }
    var v10 string = ("foo bar" + "")
    for fi11 := 0; fi11 < 3; fi11++ {
        switch ((v10 + v10) + ((v10 + v10) + "foo bar")) {
        case "hello":
            fmt.Println(fi11, v1)
        }
        var v12 int16 = -359
        switch ("\n\t" + "\n\t") {
        case "":
            fmt.Println(fi11, v10, v1)
        case "\n\t":
            v13 := GenS1{F0: 654}
            _ = v13
        case "a":
            v14 := GenS2{F0: 101}
            _ = v14
        default:
            _ = len(v5)
        }
        _ = fi11
        _ = v12
    }
    var v15 uint16 = 0
    fmt.Println(v4, v1)
    if false {
        _ = len(v2)
        for fi16 := 0; fi16 < 8; fi16++ {
            fmt.Println(v10)
            var v17 string = "foo bar"
            var v18 uint64 = 42
            _ = fi16
            _ = v17
            _ = v18
        }
        fmt.Println(v4)
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        if (0 < 371) {
            if len(v3) > 0 {
                v3[0] = (v1 >> ((42) & 63))
            }
            fmt.Println(v4, v1)
            fmt.Println(fi19, v4)
        }
        v10 = (("foo bar" + v10) + v10)
        v20 := GenS1{F0: -715}
        _ = fi19
        _ = v20
    }
    v1 = v1
    v21 := GenS3{F0: 0}
    var v22 int32 = -1
    fmt.Println(v1)
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 uint16 = (v15 * (v15 << ((255) & 63)))
        fmt.Println(v15)
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
