package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int16 = 290
    v2 := []uint32{1, 908}
    var v3 bool = true
    var v4 int = 32768
    var v5 uint8 = 127
    var v6 float32 = ((((1e10 - 0.1) + float32(0.1)) * (1.5 / (3.14159 * (((float32(0.1) * (((3.14159 / 1.5) + -2.25) / 0.1)) / float32((-2.25 + 1.5))) * 1e10)))) - 1.5)
    var v7 float64 = 0.0
    var v8 uint64 = 2
    v7 += v7
    v1 = 293
    var v9 int = 42
    _ = len(v2)
    if v3 {
        v1 = 32767
        v8 -= v8
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        if true {
            var v11 float32 = -2.25
            _ = v11
        }
        _ = fi10
    }
    switch (("a" + ((("a" + "a") + ("a" + "\n\t")) + "a")) + ("\n\t" + (("a" + ("foo bar" + "foo bar")) + (("" + "\n\t") + ("foo bar" + "\n\t"))))) {
    case "hello":
        if (!((18446744073709551615 <= 42) && v3)) {
            switch "hello" {
            case "foo bar":
                if (v3 || v3) {
                    v4 = 127
                }
            case "\n\t":
                v3 = (969 >= (v1 + v1))
            case "hello":
                var v12 string = (("a" + "") + "hello")
                _ = v12
            }
            for fi13 := 0; fi13 < 8; fi13++ {
                for fi14 := 0; fi14 < 12; fi14++ {
                    v8 -= v8
                    _ = fi14
                }
                var v15 string = ""
                _ = fi13
                _ = v15
            }
            switch "a" {
            case "a":
                var v16 int32 = 2147483647
                _ = v16
            case "hello":
                v17 := GenS1{F0: 42}
                _ = v17
            case "":
                for fi18 := 0; fi18 < 13; fi18++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v9 = v4
                    if v3 {
                        var v19 int32 = 42
                        _ = v19
                    }
                    _ = fi18
                }
            }
        }
    case "a":
        v9 -= (-(v9))
    case "\n\t":
        for fi20 := 0; fi20 < 14; fi20++ {
            if len(v2) > 0 {
                v2[0] = 42
            }
            fmt.Println(v5, v1)
            var v21 int32 = -2147483648
            _ = fi20
            _ = v21
        }
    }
    var v22 string = "hello"
    v23 := GenS1{F0: (v8 >> ((121) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
