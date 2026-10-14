package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint8) int64 {
    p0 = p0
    var v6 int8 = 127
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 uint32 = 351
    var v2 uint64 = 1
    v3 := make(map[int]uint32)
    v4 := []int{-959}
    v5 := []string{("" + "foo bar"), (((((("hello" + "") + "hello") + "\n\t") + "\n\t") + ("\n\t" + "")) + ""), ""}
    var v7 int8 = 100
    if true {
        if false {
            if true {
                v8 := GenS1{F0: 2147483647}
                v7 -= (v7 % ((v7) | 1))
                var v9 string = (("\n\t" + "foo bar") + "foo bar")
                _ = v8
                _ = v9
            }
            if false {
                v2 = (v2 + (v2 * (v2 >> ((uint(v2)) & 63))))
            } else {
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 uint16 = 65535
                    v1 += (^(v1))
                    v7 = (v7 >> ((897) & 63))
                    _ = fi10
                    _ = v11
                }
                for fi12 := 0; fi12 < 4; fi12++ {
                    if true {
                        for fi13 := 0; fi13 < 10; fi13++ {
                            switch ("hello" + (((("foo bar" + "a") + "") + "") + "")) {
                            case "hello":
                                var v14 int = 32768
                                _ = v14
                            }
                            fmt.Println(fi12)
                            v7 -= v7
                            _ = fi13
                        }
                        var v15 float64 = -2.25
                        _ = v15
                    }
                    _ = fi12
                }
                var v16 uint16 = 2
                _ = v16
            }
            var v17 uint16 = 0
            _ = v17
        } else {
            v18 := GenS2{F0: v2}
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 string = ("\n\t" + "hello")
                for fi21 := 0; fi21 < 11; fi21++ {
                    v2 -= 18446744073709551615
                    fmt.Println(v20)
                    _ = fi21
                }
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        fmt.Println(v2)
    } else {
        fmt.Println(v7, v7, v2)
    }
    _ = len(v3)
    var v22 bool = (2 < 2)
    v1 -= (^(v1))
    v22 = v22
    v7 = v7
    var v23 string = "hello"
    v24 := GenS1{F0: 0}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v25 int16 = -627
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
