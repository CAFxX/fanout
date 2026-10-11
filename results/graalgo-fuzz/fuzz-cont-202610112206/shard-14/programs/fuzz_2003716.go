package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
func genFn0(p0 int, p1 uint64) int16 {
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 int32 = 2147483647
        switch ("hello" + "foo bar") {
        case "":
            switch (("\n\t" + "foo bar") + "foo bar") {
            case "hello":
                fmt.Println(fi5, fi5, p0)
            case "a":
                p0 = 4294967295
            case "foo bar":
                p1 = p1
            }
        case "foo bar":
            v7 := []int{fi5, fi5, 2147483647, fi5}
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    if ((-2.25 / 1e10) >= 1.5) {
        var v8 int32 = -2147483648
        _ = v8
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 float32 = (1e10 + -2.25)
    var v2 int32 = 824
    var v3 uint = 255
    var v4 uint8 = 2
    switch "a" {
    case "\n\t":
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint16 = 0
            _ = fi9
            _ = v10
        }
    case "hello":
        v2 -= v2
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if false {
            v12 := GenS1{F0: 291}
            for fi13 := 0; fi13 < 10; fi13++ {
                switch "\n\t" {
                case "":
                    if (true && false) {
                        if true {
                            v14 := GenS1{F0: -1}
                            _ = v14
                        }
                        v1 -= v1
                        var v15 string = ("a" + "foo bar")
                        _ = v15
                    }
                }
                var v16 uint = v3
                _ = fi13
                _ = v16
            }
            _ = v12
        }
        switch "a" {
        case "a":
            for fi17 := 0; fi17 < 20; fi17++ {
                switch ((((("a" + "a") + "") + "") + ("\n\t" + "\n\t")) + "hello") {
                case "hello":
                    fmt.Println(v1, fi17)
                case "foo bar":
                    fmt.Println(v3, v3)
                case "a":
                    fmt.Println(fi11, v1)
                }
                _ = fi17
            }
        default:
            if false {
                v18 := GenS1{F0: fi11}
                var v19 float64 = float64(((v1 / -2.25) - v1))
                v20 := GenS2{F0: "\n\t"}
                _ = v18
                _ = v19
                _ = v20
            } else {
                var v21 int8 = 127
                _ = v21
            }
        }
        var v22 uint64 = 1
        _ = fi11
        _ = v22
    }
    var v23 int32 = int32(v4)
    if true {
        if false {
            v3 -= (^(v3))
            v4 += (^(v4))
        } else {
            fmt.Println(v23, v1, v3)
            v1 = 3.14159
        }
        var v24 float64 = 1.5
        fmt.Println(v4)
        _ = v24
    } else {
        var v25 float64 = (0.0 / 3.14159)
        _ = v25
    }
    var v26 int = 128
    var v27 int32 = (v2 >> ((v3) & 63))
    fmt.Println(v3, v27)
    var v28 int32 = 1
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 uint32 = 2
        _ = fi29
        _ = v30
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
