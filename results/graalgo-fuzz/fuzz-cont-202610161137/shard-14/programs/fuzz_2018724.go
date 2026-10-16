package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int8) uint8 {
    fmt.Println(p0)
    v6 := []int{}
    _ = p0
    _ = v6
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 2
    var v3 uint16 = 2
    var v4 float64 = 3.14159
    var v5 int64 = -104
    if (true || (42 == 818)) {
        if ((!true) || false) {
            var v7 int16 = 1000
            var v8 uint64 = 42
            var v9 int8 = 27
            _ = v7
            _ = v8
            _ = v9
        }
        if (127 >= v2) {
            v10 := GenS1{F0: -32768}
            v11 := GenS1{F0: 1}
            for fi12 := 0; fi12 < 9; fi12++ {
                fmt.Println(fi12)
                v3 += (v3 + 689)
                _ = fi12
            }
            _ = v10
            _ = v11
        } else {
            if ((0 != -986) || ("\n\t" >= "")) {
                var v13 uint = 0
                _ = v13
            } else {
                switch 32768 {
                case 400:
                    var v14 uint32 = 4294967295
                    _ = v14
                case 1:
                    fmt.Println(v5)
                case -101:
                    v15 := GenS1{F0: 1000}
                    _ = v15
                }
            }
            v16 := []int{}
            _ = v16
        }
    } else {
        v2 += v2
        var v17 uint64 = 42
        _ = v17
    }
    var v18 uint64 = 42
    if false {
        var v19 string = "hello"
        v18 = (v18 - v18)
        if ((v3 / ((v3) | 1)) >= 32767) {
            v19 = v19
        } else {
            _ = len(v1)
            var v20 int16 = -1
            v21 := []int{}
            _ = v20
            _ = v21
        }
        _ = v19
    } else {
        fmt.Println(v3, v18, v4)
        if true {
            v2 = (v2 & genFn0(-1))
        }
    }
    var v22 int16 = -32768
    if (65535 <= 491) {
        v2 += v2
    }
    fmt.Println(v2, v5)
    v4 = v4
    var v23 uint64 = (v18 * v18)
    fmt.Println(v3, v23, v5)
    v4 = (0.0 - v4)
    fmt.Println(v2, v3)
    var v24 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
