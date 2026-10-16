package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
func genFn0() int64 {
    switch (("hello" + "a") + "") {
    case "hello":
        v7 := GenS1{F0: 65535}
        _ = v7
    default:
        var v8 float32 = -2.25
        _ = v8
    }
    var v9 int16 = 0
    var v10 int64 = -1
    switch (("a" + "\n\t") + "foo bar") {
    case "foo bar":
        v11 := GenS1{F0: 686}
        _ = v11
    default:
        var v12 int = 42
        _ = v12
    }
    var v13 float32 = (3.14159 - ((1.5 * (0.1 * (0.1 + 1.5))) * 1e10))
    _ = v10
    _ = v13
    _ = v9
    return v10
}
func genFn1(p0 int32) string {
    var v14 float64 = (1e10 / 1.0)
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            if true {
                fmt.Println(v14)
                v17 := GenS2{F0: (91 <= 0)}
                switch ("a" + "a") {
                case "a":
                    switch (fi16 / ((fi15) | 1)) {
                    case 4294967295:
                        v14 += v14
                    case 32768:
                        p0 = (p0 >> ((1) & 63))
                    case 0:
                        fmt.Println(fi16, p0, fi15)
                    default:
                        v18 := GenS1{F0: 185}
                        _ = v18
                    }
                }
                _ = v17
            }
            var v19 int16 = -1
            v19 = (v19 >> ((1) & 63))
            var v20 float64 = v14
            _ = fi16
            _ = v19
            _ = v20
        }
        _ = fi15
    }
    if false {
        if false {
            fmt.Println(p0, p0)
            switch ("\n\t" + "") {
            case "a":
                fmt.Println(p0, p0)
            default:
                fmt.Println(p0, p0)
            }
        }
        for fi21 := 0; fi21 < 9; fi21++ {
            v14 = v14
            _ = fi21
        }
        var v22 int16 = -1
        _ = v22
    }
    var v23 bool = (0.0 == (3.14159 - (1e10 - 1.5)))
    _ = p0
    _ = v14
    _ = v23
    return "hello"
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 2147483648
    var v3 uint64 = 1
    var v4 float32 = (3.14159 * (3.14159 - ((1.5 * (3.14159 - 3.14159)) - 0.0)))
    v5 := []string{}
    v6 := make(map[int]string)
    var v24 int16 = -1
    var v25 uint32 = v2
    v4 -= ((v4 + (3.14159 / 1.5)) - float32(((float64(3.14159) * ((float64(0.1) / 1e10) - (0.1 * 1e10))) * (float64((v4 * v4)) / (-2.25 - 1.5)))))
    v3 = v3
    fmt.Println(v24, v4, v4)
    if (2147483647 != -2147483648) {
        if len(v5) > 0 {
            v5[0] = "a"
        }
        if false {
            _ = len(v6)
        }
    }
    var v26 uint16 = 65535
    fmt.Println(v25)
    v27 := GenS2{F0: ("hello" <= "hello")}
    var v28 uint32 = 892
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
