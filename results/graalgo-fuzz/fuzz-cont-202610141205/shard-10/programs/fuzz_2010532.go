package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]bool)
    var v3 int64 = 42
    var v4 uint64 = 0
    v5 := make(map[int]string)
    var v6 string = "hello"
    fmt.Println(v1, v6, v6)
    if ((v6 + v6) >= "hello") {
        var v7 int32 = 136
        v1 -= 1000
        var v8 string = v6
        _ = v7
        _ = v8
    }
    v6 = v6
    var v9 int64 = (v3 + 42)
    v6 = "hello"
    if (!true) {
        v3 += v3
        for fi10 := 0; fi10 < 9; fi10++ {
            fmt.Println(v4)
            _ = fi10
        }
        v4 = v4
    } else {
        v3 = (^(v9))
    }
    v3 = (v9 & v3)
    fmt.Println(v4, v4)
    var v11 int8 = 125
    if true {
        v11 = (v11 % (((v11 + v11)) | 1))
        var v12 string = ("a" + (("a" + "\n\t") + ("\n\t" + (((("" + v6) + ("" + ("hello" + "hello"))) + ("\n\t" + v6)) + v6))))
        var v13 int16 = (^(v1))
        _ = v12
        _ = v13
    } else {
        if (440 < 2) {
            var v14 uint8 = 128
            fmt.Println(v9, v14)
            v1 = (v1 + v1)
            _ = v14
        }
        var v15 string = "foo bar"
        _ = v15
    }
    var v16 float64 = ((-2.25 - 0.0) + (3.14159 * -2.25))
    var v17 int8 = v11
    v2[int(v11)] = (!false)
    fmt.Println(v3, v9, v16)
    v16 = 1.5
    v17 += v17
    switch (("\n\t" + v6) + (((("\n\t" + ((v6 + ((v6 + v6) + v6)) + ((((v6 + v6) + ("hello" + v6)) + v6) + "foo bar"))) + (v6 + v6)) + "\n\t") + (v6 + v6))) {
    case "\n\t":
        for fi18 := 0; fi18 < 9; fi18++ {
            v9 += v9
            v19 := GenS1{F0: (v17 == -79)}
            _ = fi18
            _ = v19
        }
    case "a":
        v17 -= v11
    }
    v16 = (((0.1 / v16) / v16) / ((((((3.14159 * v16) / float64((float32((((v16 / v16) * v16) / (-2.25 + (v16 * (float64((0.1 * (float32(1.5) / 1e10))) * v16))))) - 1.5))) + (v16 / v16)) * 1e10) * float64(0.0)) / float64(((float32(float64(0.1)) * (1.5 / 1.0)) * (1e10 * 3.14159)))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
