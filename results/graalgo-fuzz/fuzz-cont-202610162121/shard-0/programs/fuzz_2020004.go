package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint) string {
    switch -835 {
    case -9223372036854775808:
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 float32 = 0.1
            v9 := []int{fi7}
            if (255 >= 42) {
                _ = len(v9)
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return "\n\t"
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]string)
    var v3 int32 = 2147483647
    v4 := []uint32{654, 42, 65535}
    var v5 float32 = (1.5 + float32(1.5))
    var v6 string = ("" + ("\n\t" + (("\n\t" + ("" + "hello")) + "foo bar")))
    _ = len(v2)
    v10 := GenS1{F0: (^(v1)), F1: -1}
    if true {
        fmt.Println(v6, v6)
        var v11 bool = (42 <= -9223372036854775808)
        _ = v11
    }
    var v12 uint32 = 62
    v12 += v12
    v13 := GenS2{F0: 9223372036854775807}
    if (v5 >= 1.5) {
        switch v6 {
        case "foo bar":
            for fi14 := 0; fi14 < 8; fi14++ {
                if (0 > 570) {
                    v3 -= v3
                    v2[v6] = v6
                } else {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v1 += (v1 * v1)
                        v5 = 1.5
                        switch ((v6 + (v6 + (("hello" + ((v6 + "a") + v6)) + (v6 + v6)))) + v6) {
                        case "":
                            fmt.Println(fi14)
                        }
                        _ = fi15
                    }
                    if false {
                        v12 -= 2147483648
                        for fi16 := 0; fi16 < 4; fi16++ {
                            var v17 string = "\n\t"
                            _ = fi16
                            _ = v17
                        }
                    }
                }
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 string = v6
                    _ = fi18
                    _ = v19
                }
                _ = fi14
            }
        case "\n\t":
            v12 += (v12 << ((632) & 63))
        }
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        if true {
            if (!(1 >= 937)) {
                v3 -= (v3 << ((v1) & 63))
                v1 = v1
            }
        } else {
            v21 := GenS2{F0: -1}
            _ = len(v4)
            _ = v21
        }
        var v22 int16 = 0
        _ = fi20
        _ = v22
    }
    v6 = ((v6 + v6) + "a")
    var v23 int32 = v3
    v6 = ((v6 + ("hello" + "hello")) + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
