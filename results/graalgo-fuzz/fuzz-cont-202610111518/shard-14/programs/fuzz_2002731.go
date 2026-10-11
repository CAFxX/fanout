package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 bool, p1 uint, p2 uint32) uint64 {
    v5 := GenS1{F0: 1}
    if (-715 == 0) {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 bool = p0
            var v8 int16 = -105
            var v9 uint8 = 2
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    }
    v10 := []int{1, 4294967295, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    return 881
}

func main() {
    var v1 uint64 = 1
    var v2 int32 = 1
    v3 := make(map[string]bool)
    var v4 bool = false
    v4 = (((105 < -128) && true) && v4)
    v4 = (225 != 128)
    if v4 {
        if (0 <= -594) {
            var v11 uint32 = 510
            fmt.Println(v11)
            if (true && true) {
                v1 += v1
            } else {
                var v12 string = ""
                v2 = v2
                v12 = (v12 + ((((v12 + "hello") + v12) + v12) + v12))
                _ = v12
            }
            _ = v11
        } else {
            switch "hello" {
            case "\n\t":
                for fi13 := 0; fi13 < 13; fi13++ {
                    if true {
                        v2 -= -1
                        fmt.Println(fi13)
                    } else {
                        if (42 == -467) {
                            v14 := []int{fi13}
                            _ = v14
                        } else {
                            fmt.Println(v4, v4, fi13)
                            var v15 int = fi13
                            _ = v15
                        }
                    }
                    _ = fi13
                }
            case "a":
                v2 = 42
            }
            v2 = v2
            v16 := []int{457, 128}
            _ = v16
        }
    }
    v2 -= (v2 >> ((0) & 63))
    for fi17 := 0; fi17 < 3; fi17++ {
        v4 = (!true)
        v18 := GenS1{F0: v2}
        v4 = false
        _ = fi17
        _ = v18
    }
    v2 -= (v2 | v2)
    if (0.1 != (-2.25 / 1.0)) {
        v2 -= v2
    }
    v19 := GenS1{F0: v2}
    var v20 bool = (-128 == 1)
    v20 = false
    if (!v20) {
        if v4 {
            if (373 == 917) {
                v20 = (1 < 255)
                var v21 float64 = 1.5
                if true {
                    v3[("\n\t" + ("\n\t" + "foo bar"))] = v4
                }
                _ = v21
            } else {
                var v22 float32 = -2.25
                v23 := GenS1{F0: v2}
                _ = v22
                _ = v23
            }
            var v24 uint8 = 0
            v2 = v2
            _ = v24
        }
    }
    fmt.Println(v4, v4, v1)
    v20 = (!(("hello" + ("foo bar" + "\n\t")) > ("foo bar" + "foo bar")))
    fmt.Println(v1, v20, v4)
    var v25 int64 = 42
    fmt.Println(v2, v4)
    var v26 uint32 = 464
    var v27 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
