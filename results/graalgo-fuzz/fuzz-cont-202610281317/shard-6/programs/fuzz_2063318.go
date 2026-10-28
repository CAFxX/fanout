package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 string; F2 bool }

func main() {
    v1 := make(map[string]string)
    v2 := []int{127, -604, 80, -1}
    var v3 uint8 = 128
    var v4 int16 = 575
    var v5 int = 9223372036854775807
    v6 := []string{"\n\t", ("" + ("hello" + ("\n\t" + "a"))), "", "hello"}
    v3 += (v3 ^ 42)
    var v7 string = "\n\t"
    delete(v1, ((("\n\t" + v7) + (v7 + ((v7 + (("a" + v7) + "")) + v7))) + "\n\t"))
    var v8 int = v5
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v9 string = v7
    var v10 int16 = v4
    v3 = v3
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS1{F0: 797}
            var v14 bool = true
            v8 -= (-(fi11))
            fmt.Println(v5, v3, v8)
            if ((-(v10)) == (v4 >> ((1) & 63))) {
                v8 = -1
                if (32767 == 187) {
                    v15 := GenS1{F0: 42}
                    _ = v15
                } else {
                    v16 := GenS2{F0: 893, F1: (v7 + ("hello" + v9)), F2: v14}
                    _ = v16
                }
                v8 += (v8 & -2147483648)
            } else {
                var v17 string = v9
                var v18 uint = 1
                v14 = v14
                _ = v17
                _ = v18
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        var v19 float32 = ((3.14159 * (1e10 - (0.0 + 0.1))) - ((1e10 / ((3.14159 + (0.1 - (0.1 * 1.5))) / 1.5)) + 1e10))
        _ = fi11
        _ = v19
    }
    var v20 int32 = 42
    for fi21 := 0; fi21 < 17; fi21++ {
        if true {
            if (0.1 >= 1e10) {
                v10 = v10
                var v22 string = v9
                _ = v22
            }
        }
        var v23 int = 128
        _ = fi21
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
