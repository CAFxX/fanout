package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
func genFn0() int16 {
    v7 := GenS1{F0: 640}
    var v8 float32 = (1e10 / 3.14159)
    v8 -= (float32(-2.25) / v8)
    v9 := []int{32767, 128, -2147483648}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int32 = -349
    var v2 int8 = 127
    v3 := []string{"", (((((((("a" + "foo bar") + ("hello" + "a")) + (((("hello" + ("\n\t" + "a")) + "\n\t") + "\n\t") + "\n\t")) + (((("" + "\n\t") + "") + ("" + ((("hello" + "foo bar") + "a") + ""))) + "")) + (("\n\t" + ("a" + (("hello" + "\n\t") + (("foo bar" + "\n\t") + (("hello" + "a") + ("hello" + "foo bar")))))) + "")) + "hello") + "a") + "")}
    var v4 int64 = -73
    var v5 uint8 = 1
    var v6 int8 = v2
    var v10 int16 = 32767
    var v11 uint64 = 342
    var v12 uint16 = 65535
    v12 = 754
    var v13 float32 = 3.14159
    var v14 float64 = ((-2.25 - 1.5) + 3.14159)
    var v15 int16 = v10
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int16 = 1000
        for fi18 := 0; fi18 < 12; fi18++ {
            v11 -= (v11 & 18446744073709551615)
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v12)
        _ = fi19
    }
    if true {
        v20 := GenS1{F0: 0}
        switch "a" {
        case "hello":
            v5 = (v5 & 65)
        case "foo bar":
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 float32 = float32(v14)
                _ = fi21
                _ = v22
            }
        case "":
            fmt.Println(v13, v13, v15)
        }
        _ = v20
    }
    v3 = append(v3, "a")
    if true {
        v23 := GenS1{F0: v12}
        v24 := GenS2{F0: ""}
        v14 = v14
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: v12}
    v10 = v10
    v1 += v1
    v14 = -2.25
    if ((v15 >> ((503) & 63)) <= (v10 << ((106) & 63))) {
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    } else {
        v11 += 185
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
