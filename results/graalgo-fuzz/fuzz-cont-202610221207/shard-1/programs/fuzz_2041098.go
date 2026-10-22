package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 string; F2 uint8 }
func genFn0(p0 uint32, p1 uint64, p2 int) int32 {
    p2 = (p2 * p2)
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: -2147483648}
        var v6 uint32 = (p0 >> ((0) & 63))
        if true {
            if (!(uint32(p1) >= v6)) {
                var v7 uint8 = 127
                _ = v7
            } else {
                if ((fi4 + fi4) == (p2 | fi4)) {
                    switch fi4 {
                    case -738:
                        v6 -= (^(p0))
                    case -2147483648:
                        switch "hello" {
                        case "":
                            var v8 string = ("" + "\n\t")
                            _ = v8
                        case "foo bar":
                            p2 -= 32768
                        }
                    case 2147483647:
                        var v9 uint64 = p1
                        _ = v9
                    default:
                        var v10 int = 4294967295
                        _ = v10
                    }
                    v11 := GenS2{F0: 42, F1: "\n\t", F2: 224}
                    if (1 != 42) {
                        var v12 int16 = -32768
                        _ = v12
                    }
                    _ = v11
                }
                v13 := []int{fi4, (-(p2))}
                _ = v13
            }
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    var v14 int8 = -35
    var v15 int8 = -78
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return -2147483648
}

func main() {
    var v1 int8 = -1
    var v2 uint = 98
    var v3 float32 = (0.1 / 1e10)
    var v16 string = ("hello" + "\n\t")
    switch 4294967295 {
    case 2147483647:
        var v17 uint16 = 255
        _ = v17
    case 370:
        v16 = (v16 + ((("a" + "\n\t") + v16) + ("a" + ("hello" + v16))))
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v16 = v16
        _ = fi18
    }
    if false {
        var v19 uint16 = 1
        _ = v19
    }
    var v20 uint16 = 255
    v16 = "foo bar"
    v16 = ""
    v21 := GenS1{F0: -2147483648}
    var v22 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
