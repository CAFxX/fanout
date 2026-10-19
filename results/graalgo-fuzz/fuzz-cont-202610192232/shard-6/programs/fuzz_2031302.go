package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint64; F2 bool }
type GenS3 struct { F0 uint8; F1 uint }

func main() {
    var v1 uint16 = 42
    var v2 string = ""
    v3 := []uint32{1, 2147483648}
    var v4 int64 = -1
    v5 := []string{v2, (v2 + v2), ("hello" + (v2 + "foo bar"))}
    var v6 bool = true
    var v7 uint16 = (v1 / ((v1) | 1))
    var v8 bool = false
    v9 := GenS1{F0: -682}
    for fi10 := 0; fi10 < 15; fi10++ {
        switch (fi10 | (fi10 | (fi10 / (((-(fi10))) | 1)))) {
        case 4294967295:
            var v11 string = "hello"
            _ = v11
        }
        _ = fi10
    }
    fmt.Println(v7)
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 uint16 = v7
        var v14 int16 = 32767
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v15 := GenS2{F0: (!v6), F1: 4294967295, F2: v6}
    if len(v3) > 0 {
        v3[0] = 65535
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 float32 = (float32((0.0 - float64((3.14159 - (float32((float64(((((float32((float64(-2.25) - (float64(0.0) / 1.5))) / 3.14159) / 1e10) + 0.0) * 1.5)) + ((float64(0.0) / (float64((-2.25 - 3.14159)) * float64(0.1))) - 0.0))) / 1.0))))) - 0.0)
        v5 = append(v5, "foo bar")
        fmt.Println(v17)
        _ = fi16
        _ = v17
    }
    switch -2147483648 {
    case -2147483648:
        for fi18 := 0; fi18 < 7; fi18++ {
            fmt.Println(v8)
            _ = fi18
        }
    case 127:
        switch "\n\t" {
        case "":
            v19 := GenS3{F0: 123, F1: 42}
            _ = v19
        case "a":
            switch 32768 {
            case 127:
                v2 = v2
            }
        default:
            fmt.Println(v7)
        }
    case 42:
        v6 = (!(255 <= 0))
    }
    v2 = v2
    fmt.Println(v7, v4)
    v8 = (2147483648 == 0)
    fmt.Println(v7)
    v6 = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
