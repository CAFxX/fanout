package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 string; F2 bool }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 uint16 = 768
    var v2 uint = 911
    var v3 float64 = 1.5
    v4 := make(map[string]int)
    var v5 uint8 = 255
    v6 := GenS1{F0: ("a" + "a"), F1: ((584 <= 0) && false)}
    v7 := GenS2{F0: (("a" + ("hello" + "foo bar")) + ("foo bar" + "hello"))}
    v8 := GenS2{F0: "a"}
    v3 -= v3
    v2 -= (v2 ^ 65535)
    if false {
        var v9 bool = (127 >= -1)
        var v10 bool = v9
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 14; fi11++ {
            if (!false) {
                v12 := GenS1{F0: ("" + "hello"), F1: true}
                _ = v12
            } else {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := []int{fi13, int(v2), (fi13 & fi11)}
                    _ = fi13
                    _ = v14
                }
                v15 := []int{}
                _ = v15
            }
            var v16 int = fi11
            _ = fi11
            _ = v16
        }
        if false {
            var v17 uint8 = v5
            _ = v17
        } else {
            if ("" >= "foo bar") {
                delete(v4, "hello")
                var v18 int8 = 127
                _ = v18
            } else {
                if ((-128 > 0) || (v5 > v5)) {
                    var v19 uint32 = 161
                    v3 += 1.5
                    if false {
                        v1 += 539
                    }
                    _ = v19
                }
                v20 := []int{2147483647, 127, -2147483648, -1}
                _ = v20
            }
        }
    }
    v21 := GenS3{F0: v1, F1: ("foo bar" + "\n\t"), F2: false}
    v5 = 127
    var v22 int = 696
    v23 := GenS4{F0: v5}
    var v24 int32 = 615
    for fi25 := 0; fi25 < 18; fi25++ {
        fmt.Println(v22)
        _ = len(v4)
        _ = fi25
    }
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
