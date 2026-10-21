package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int32) uint32 {
    switch -1 {
    case 0:
        p0 = p0
    case 42:
        var v6 uint64 = 4294967295
        _ = v6
    case 4294967295:
        fmt.Println(p0)
    }
    var v7 bool = (!true)
    v7 = (!v7)
    _ = p0
    _ = v7
    return 2147483648
}

func main() {
    var v1 int = 2147483647
    v2 := []string{(("foo bar" + "\n\t") + ("a" + "hello")), "hello", ("hello" + (((((("hello" + "") + (((("hello" + "a") + "") + ("foo bar" + "foo bar")) + ("a" + "\n\t"))) + "\n\t") + "hello") + ("a" + "a")) + "hello")), ""}
    v3 := make(map[string]string)
    var v4 uint32 = 2147483647
    var v5 int = v1
    v4 = (v4 / (((v4 * (v4 ^ 65535))) | 1))
    switch (^(v5)) {
    case 42:
        v5 = v1
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint16 = 255
        v10 := GenS1{F0: -2147483648}
        _ = fi8
        _ = v10
        _ = v9
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v5 += v5
    switch ("hello" + "\n\t") {
    case "":
        fmt.Println(v1)
    default:
        fmt.Println(v5, v5)
    }
    var v11 int8 = 100
    var v12 uint8 = 72
    v3["a"] = "hello"
    switch (((((("foo bar" + "a") + "a") + ("hello" + "\n\t")) + (("foo bar" + ("hello" + ("\n\t" + ("a" + "\n\t")))) + "")) + "a") + "") {
    case "foo bar":
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 float32 = float32(1e10)
            _ = fi13
            _ = v14
        }
    case "hello":
        v2 = append(v2, ("a" + "\n\t"))
    case "a":
        var v15 uint64 = 42
        _ = v15
    }
    fmt.Println(v12)
    for fi16 := 0; fi16 < 14; fi16++ {
        switch v5 {
        case -143:
            switch (fi16 / ((v5) | 1)) {
            case 579:
                v12 -= 127
            }
        case 1:
            fmt.Println(fi16, v4)
        case 853:
            if false {
                for fi17 := 0; fi17 < 10; fi17++ {
                    if (!true) {
                        v18 := v3[("a" + ("\n\t" + ""))]
                        _ = v18
                    }
                    _ = fi17
                }
            }
        default:
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 uint64 = 18446744073709551615
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if false {
                        v20 = (v20 % ((0) | 1))
                    }
                    var v21 int16 = 42
                    _ = v21
                } else {
                    fmt.Println(v11, v20)
                }
                if false {
                    var v22 uint = 1
                    var v23 uint = v22
                    v24 := GenS1{F0: int32(v1)}
                    var v25 float64 = (-2.25 / 3.14159)
                    var v26 int = v1
                    _ = v22
                    _ = v23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                _ = fi19
                _ = v20
            }
        }
        _ = fi16
    }
    switch "hello" {
    case "\n\t":
        fmt.Println(v12)
    default:
        v11 = (v11 + v11)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
