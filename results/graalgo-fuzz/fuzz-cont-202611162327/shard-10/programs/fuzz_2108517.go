package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int = 770
    var v3 string = "hello"
    var v4 uint16 = 32767
    if (!(float32(1.5) <= (3.14159 - float32((3.14159 - 0.0))))) {
        fmt.Println(v4, v1, v3)
        v5 := []int{(^(v2))}
        if (!true) {
            v4 = (v4 - v4)
        } else {
            v5 = append(v5, (-(v2)))
            if true {
                var v6 float64 = float64(0.0)
                _ = v6
            } else {
                switch (v3 + "\n\t") {
                case "":
                    v1 = (^(v1))
                }
                v7 := GenS1{F0: (!(!(0 <= int32(v1))))}
                v3 = "foo bar"
                _ = v7
            }
        }
        _ = v5
    }
    var v8 string = "a"
    v9 := []int{(-(v2)), v2}
    switch 2147483647 {
    case 2147483647:
        for fi10 := 0; fi10 < 6; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 bool = true
                _ = fi11
                _ = v12
            }
            v8 = v8
            v3 = ((v8 + (v3 + v8)) + (("\n\t" + "hello") + "hello"))
            _ = fi10
        }
    default:
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint32 = 253
            var v15 string = v3
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v8 = (((((((v3 + v8) + "foo bar") + "hello") + v8) + ("\n\t" + (("foo bar" + ("\n\t" + v8)) + v8))) + (v3 + v3)) + ("foo bar" + (v8 + v8)))
    for fi16 := 0; fi16 < 10; fi16++ {
        if false {
            v9 = append(v9, 32768)
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v17 uint16 = v4
            _ = v17
        }
        switch (fi16 - (^(fi16))) {
        case -2147483648:
            v1 -= v1
        case -537:
            v9 = append(v9, (-(v2)))
        }
        _ = fi16
    }
    switch (v8 + v8) {
    case "\n\t":
        var v18 float32 = (0.1 / 1.0)
        _ = v18
    case "foo bar":
        var v19 bool = true
        _ = v19
    case "hello":
        var v20 int64 = v1
        _ = v20
    default:
        if (uint64(v4) == 845) {
            v3 = "hello"
            if true {
                var v21 int32 = 2147483647
                if false {
                    v9 = append(v9, (v2 >> ((255) & 63)))
                }
                var v22 string = "hello"
                _ = v21
                _ = v22
            } else {
                v8 = ("a" + "a")
            }
            var v23 int32 = 0
            _ = v23
        }
    }
    v24 := GenS1{F0: (false || true)}
    var v25 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
