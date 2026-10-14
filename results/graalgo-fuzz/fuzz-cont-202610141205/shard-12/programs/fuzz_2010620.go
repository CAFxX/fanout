package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 int64 = 0
    var v4 uint32 = 0
    var v5 float64 = (-2.25 + (1e10 + ((1e10 / (1.5 * 3.14159)) * ((0.0 * 1e10) - 1.5))))
    var v6 int = 32767
    var v7 uint8 = 128
    v6 = v6
    v6 += 749
    v5 = (float64(0.0) + v5)
    var v8 bool = (-128 < 42)
    if v8 {
        v3 += (v3 >> ((265) & 63))
        if (!(1 >= 42)) {
            switch (v6 + (v6 / ((v6) | 1))) {
            case 2147483647:
                var v9 string = "\n\t"
                _ = v9
            case 944:
                v8 = (!false)
            }
        } else {
            v10 := GenS1{F0: v7}
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 float64 = (0.1 * 0.1)
                for fi13 := 0; fi13 < 9; fi13++ {
                    v12 += 1e10
                    var v14 bool = v8
                    _ = fi13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
            var v15 string = "foo bar"
            _ = v10
            _ = v15
        }
        if ((v7 | 128) == v7) {
            var v16 int64 = v3
            v17 := []int{v6}
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 bool = ((v8 || true) || v8)
                v20 := GenS1{F0: 1}
                if len(v17) > 0 {
                    v17[0] = fi18
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v16
            _ = v17
        } else {
            var v21 int8 = 100
            v8 = (false || (v8 || v8))
            var v22 int = v6
            _ = v21
            _ = v22
        }
    }
    var v23 int16 = 0
    v23 += (^(v23))
    fmt.Println(v4, v7, v3)
    fmt.Println(v5)
    var v24 uint = 1
    var v25 int = (v6 % ((v6) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
