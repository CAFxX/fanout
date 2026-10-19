package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int32, p1 uint) int16 {
    v6 := []int{4294967295, 42, 1, 128}
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(fi7, p0)
        for fi8 := 0; fi8 < 16; fi8++ {
            if (475 > 1) {
                switch (^(fi8)) {
                case -1:
                    if true {
                        p0 = p0
                        if len(v6) > 0 {
                            v6[0] = -9223372036854775808
                        }
                    }
                }
            } else {
                var v9 int32 = p0
                var v10 int32 = p0
                _ = v10
                _ = v9
            }
            _ = fi8
        }
        _ = fi7
    }
    v6 = append(v6, -178)
    _ = p0
    _ = p1
    _ = v6
    return 32767
}

func main() {
    var v1 bool = false
    var v2 int = 2147483647
    var v3 string = ("hello" + ("\n\t" + "\n\t"))
    v4 := make(map[int]bool)
    v5 := make(map[int]string)
    switch "foo bar" {
    case "foo bar":
        if (float32((1e10 + 1e10)) <= 1e10) {
            var v11 int = v2
            var v12 int = (-(v2))
            if false {
                for fi13 := 0; fi13 < 8; fi13++ {
                    v14 := []int{v2, v2, 32768, -9223372036854775808}
                    if (-9223372036854775808 <= -9223372036854775808) {
                        v12 += v2
                    }
                    _ = fi13
                    _ = v14
                }
            } else {
                var v15 int16 = genFn0(2147483647, 1)
                var v16 uint16 = 682
                v17 := GenS1{F0: (v15 - v15)}
                _ = v15
                _ = v16
                _ = v17
            }
            _ = v11
            _ = v12
        }
    }
    v1 = false
    v2 -= (^(v2))
    v18 := v5[32768]
    v4[9223372036854775807] = false
    v3 = (("" + ((v18 + v3) + ((("hello" + v18) + v18) + (("foo bar" + (v3 + "hello")) + v18)))) + v3)
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 uint = 667
        _ = fi19
        _ = v20
    }
    switch (v18 + (("\n\t" + v18) + "foo bar")) {
    case "foo bar":
        v5[v2] = (v18 + "a")
    case "hello":
        var v21 uint16 = 1
        _ = v21
    case "\n\t":
        var v22 float32 = 0.0
        _ = v22
    }
    var v23 uint16 = 2
    var v24 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
