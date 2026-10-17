package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 4294967295
    var v3 float32 = 1.5
    var v4 int16 = -1
    var v5 int32 = 2147483647
    var v6 int16 = (v4 << ((255) & 63))
    var v7 int = 127
    v8 := GenS1{F0: 1}
    for fi9 := 0; fi9 < 12; fi9++ {
        v7 = 42
        var v10 string = "a"
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := []int{int(v2), (fi11 ^ (v7 << ((280) & 63)))}
        var v13 string = "\n\t"
        _ = fi11
        _ = v12
        _ = v13
    }
    v5 = v5
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS2{F0: v2}
        v16 := []int{}
        v17 := GenS3{F0: 109}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v4 = (^(v6))
        switch "\n\t" {
        case "":
            var v19 string = "hello"
            _ = v19
        case "foo bar":
            v20 := []int{fi18, (-(fi18)), (v7 | (v7 + (v7 + (fi18 >> ((880) & 63)))))}
            _ = v20
        case "hello":
            var v21 uint16 = 32767
            _ = v21
        }
        _ = fi18
    }
    v22 := GenS2{F0: v2}
    var v23 float32 = v3
    var v24 uint16 = 32767
    for fi25 := 0; fi25 < 12; fi25++ {
        fmt.Println(v4)
        fmt.Println(fi25, v2, v23)
        _ = fi25
    }
    fmt.Println(v2)
    var v26 int8 = -49
    v1 = 255
    for fi27 := 0; fi27 < 4; fi27++ {
        v28 := GenS4{F0: 663}
        _ = fi27
        _ = v28
    }
    v6 += 410
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
