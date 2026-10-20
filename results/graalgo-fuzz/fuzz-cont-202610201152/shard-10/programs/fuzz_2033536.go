package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint; F1 bool; F2 int32 }
func genFn0(p0 uint64, p1 uint64) int {
    v4 := []int{9223372036854775807, 128, 9223372036854775807}
    var v5 int64 = 42
    v5 -= v5
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int = -777
    var v2 uint = 431
    var v3 int16 = 0
    v2 += v2
    fmt.Println(v2, v2)
    v6 := GenS1{F0: 0, F1: -2147483648, F2: (v1 + v1)}
    v2 += (^(v2))
    fmt.Println(v1)
    var v7 uint8 = 15
    v2 -= v2
    var v8 int16 = 680
    v9 := GenS2{F0: v3, F1: "hello"}
    var v10 bool = true
    var v11 uint = (v2 >> ((v2) & 63))
    if (true || false) {
        if false {
            v1 += 127
            v12 := []int{v1}
            _ = v12
        } else {
            v13 := GenS1{F0: 42, F1: 0, F2: v1}
            var v14 int = genFn0(989, 4294967295)
            if (42 >= 65535) {
                switch "a" {
                case "foo bar":
                    v15 := GenS2{F0: (v8 << ((733) & 63)), F1: "a"}
                    _ = v15
                }
            }
            _ = v13
            _ = v14
        }
        v16 := GenS1{F0: v7, F1: 132, F2: (v1 ^ 1)}
        _ = v16
    }
    var v17 float32 = -2.25
    var v18 uint32 = 42
    v8 += v8
    for fi19 := 0; fi19 < 7; fi19++ {
        v20 := GenS3{F0: (v11 >> (((v11 << ((uint(v18)) & 63))) & 63)), F1: (false || (!(80 > v18))), F2: -2147483648}
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 string = "\n\t"
            v23 := []int{}
            if ((42 <= (v7 << (((v11 - 42)) & 63))) && true) {
                fmt.Println(v3, v17, v10)
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi19
        _ = v20
    }
    if v10 {
        if v10 {
            v24 := []int{v1}
            for fi25 := 0; fi25 < 10; fi25++ {
                fmt.Println(v8, v2, v1)
                fmt.Println(v17)
                v26 := GenS1{F0: (^(v7)), F1: -1, F2: (fi25 | -81)}
                _ = fi25
                _ = v26
            }
            _ = v24
        }
        for fi27 := 0; fi27 < 8; fi27++ {
            for fi28 := 0; fi28 < 15; fi28++ {
                v11 = (v11 % ((v11) | 1))
                fmt.Println(v2, v8, v18)
                var v29 int16 = (v8 | -1)
                v7 += (v7 | (v7 & (v7 ^ uint8(v18))))
                fmt.Println(fi27, v2)
                _ = fi28
                _ = v29
            }
            switch (((((("\n\t" + "hello") + (("" + ("hello" + "foo bar")) + "a")) + "") + "a") + "a") + "a") {
            case "\n\t":
                v30 := GenS1{F0: (^(v7)), F1: 1, F2: (fi27 << ((v11) & 63))}
                _ = v30
            case "":
                if false {
                    var v31 float64 = (((3.14159 / 3.14159) - 1e10) / 0.1)
                    var v32 string = "\n\t"
                    v10 = v10
                    _ = v31
                    _ = v32
                }
            }
            var v33 bool = (int(v3) != (v1 << ((v2) & 63)))
            _ = fi27
            _ = v33
        }
    } else {
        var v34 float64 = 1.5
        v35 := GenS2{F0: v3, F1: "a"}
        _ = v34
        _ = v35
    }
    v36 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
