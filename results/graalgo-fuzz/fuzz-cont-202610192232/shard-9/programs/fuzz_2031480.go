package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
func genFn0() string {
    if true {
        if true {
            var v4 int8 = -128
            _ = v4
        }
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 int16 = 32767
            _ = fi5
            _ = v6
        }
        var v7 int64 = 42
        _ = v7
    }
    var v8 uint8 = 128
    _ = v8
    return ("" + ("a" + "a"))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    v9 := GenS1{F0: 42}
    for fi10 := 0; fi10 < 1; fi10++ {
        switch 857 {
        case 141:
            var v11 int64 = 42
            _ = v11
        case 42:
            v12 := GenS1{F0: 0}
            _ = v12
        default:
            var v13 uint32 = 2147483648
            _ = v13
        }
        var v14 int = 2147483647
        v15 := GenS1{F0: 521}
        _ = fi10
        _ = v14
        _ = v15
    }
    var v16 int32 = 2147483647
    delete(v3, (("" + "") + "hello"))
    v17 := GenS2{F0: 429}
    switch ("hello" + ("hello" + ("\n\t" + ("a" + ("foo bar" + ""))))) {
    case "foo bar":
        for fi18 := 0; fi18 < 11; fi18++ {
            for fi19 := 0; fi19 < 20; fi19++ {
                fmt.Println(fi19, fi18)
                var v20 uint32 = 165
                var v21 bool = (!(false || false))
                _ = fi19
                _ = v20
                _ = v21
            }
            v22 := GenS2{F0: 0}
            v23 := GenS3{F0: 255}
            _ = fi18
            _ = v22
            _ = v23
        }
    case "a":
        var v24 int64 = 1
        _ = v24
    case "":
        if (!(76 <= -77)) {
            var v25 uint8 = 213
            _ = v25
        } else {
            var v26 float64 = (0.1 - 3.14159)
            _ = v26
        }
    default:
        v16 = 42
    }
    v16 = v16
    var v27 int32 = v16
    v16 += v27
    for fi28 := 0; fi28 < 19; fi28++ {
        v29 := []int{fi28, fi28}
        v27 = v16
        var v30 int64 = 42
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v9
}
