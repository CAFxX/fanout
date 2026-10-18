package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0(p0 uint8, p1 bool) int32 {
    p0 = (p0 * (^(p0)))
    v7 := []int{127, 687, 219, -51}
    v8 := GenS1{F0: 1, F1: 1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 int = 9223372036854775807
    var v2 int8 = -45
    var v3 int64 = 1
    var v4 float32 = 1.5
    v5 := make(map[uint32]int)
    var v6 string = ""
    v3 += v3
    v9 := GenS1{F0: -1, F1: -109}
    v4 -= 1e10
    v2 -= 100
    fmt.Println(v1)
    var v10 uint32 = 2147483647
    if false {
        switch (("a" + "hello") + ("foo bar" + v6)) {
        case "a":
            delete(v5, (^(v10)))
        default:
            switch (v6 + (v6 + (v6 + "hello"))) {
            case "hello":
                var v11 uint64 = 4294967295
                _ = v11
            case "":
                for fi12 := 0; fi12 < 7; fi12++ {
                    if ((614 > -670) && (!(!(!false)))) {
                        switch v6 {
                        case "foo bar":
                            switch (v1 & int(v10)) {
                            case 0:
                                v10 = v10
                            }
                        }
                        v4 -= (v4 * 1e10)
                        v13 := []int{fi12}
                        _ = v13
                    } else {
                        fmt.Println(v10, fi12, v2)
                    }
                    delete(v5, (v10 * (v10 & 47)))
                    var v14 string = ""
                    _ = fi12
                    _ = v14
                }
            case "foo bar":
                v6 = (v6 + "hello")
            }
        }
    }
    var v15 uint64 = 0
    fmt.Println(v3, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
