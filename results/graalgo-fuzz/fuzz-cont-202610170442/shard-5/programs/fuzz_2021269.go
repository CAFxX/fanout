package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint = 394
    var v2 float32 = ((1.5 * ((3.14159 - -2.25) / ((1.5 / 1.0) * (0.1 * (1.5 + 1.5))))) * 1.5)
    v3 := []string{"hello", ("foo bar" + ((("" + "\n\t") + ((("" + (("" + ("a" + "")) + "\n\t")) + "hello") + (("\n\t" + "foo bar") + ("\n\t" + ("foo bar" + ("hello" + "hello")))))) + "foo bar")), ("" + "\n\t")}
    var v4 float64 = 1e10
    var v5 string = ("foo bar" + "")
    var v6 float32 = 1.5
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v5 = "foo bar"
    switch 128 {
    case -1:
        if (492 < -32768) {
            var v7 uint = 128
            _ = v7
        } else {
            if len(v3) > 0 {
                v3[0] = (v5 + ("hello" + ""))
            }
        }
    case 127:
        v8 := GenS1{F0: 42}
        _ = v8
    case -2147483648:
        v4 = v4
    }
    var v9 float64 = 1e10
    if false {
        v10 := GenS1{F0: -1}
        v3 = append(v3, "hello")
        var v11 uint16 = 32767
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS1{F0: 127}
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint16 = 2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v16 int32 = 344
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = len(v3)
        _ = fi12
        _ = v13
    }
    switch -9223372036854775808 {
    case 128:
        v17 := GenS1{F0: -1}
        _ = v17
    }
    fmt.Println(v4, v6, v9)
    var v18 bool = (127 == 57)
    v6 += 1.5
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if v18 {
            fmt.Println(v1, v18)
            var v19 int32 = 1
            _ = v19
        }
        v2 += ((v6 - float32(0.1)) * (v6 + v6))
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    var v20 uint8 = 12
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
