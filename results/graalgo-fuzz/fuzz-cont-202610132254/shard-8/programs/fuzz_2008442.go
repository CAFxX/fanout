package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 string }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0(p0 string) int64 {
    switch 32767 {
    case -2147483648:
        v4 := []int{1, 177}
        _ = v4
    default:
        var v5 string = (p0 + (((((p0 + (p0 + (p0 + p0))) + ((p0 + p0) + "foo bar")) + p0) + ((("foo bar" + "a") + (p0 + (p0 + "a"))) + "hello")) + (p0 + p0)))
        _ = v5
    }
    v6 := GenS1{F0: 373, F1: 0, F2: "a"}
    if true {
        v7 := GenS1{F0: -1, F1: 793, F2: p0}
        switch p0 {
        case "\n\t":
            if true {
                switch (p0 + (p0 + ("a" + p0))) {
                case "a":
                    p0 = p0
                case "foo bar":
                    v8 := GenS1{F0: 0, F1: -2147483648, F2: p0}
                    _ = v8
                case "":
                    p0 = p0
                }
            }
        }
        _ = v7
    } else {
        p0 = ""
        p0 = p0
    }
    if false {
        fmt.Println(p0)
        if ((-1 < 678) && (false && false)) {
            switch p0 {
            case "hello":
                p0 = "foo bar"
            }
            var v9 bool = (42 > 42)
            _ = v9
        } else {
            p0 = p0
            v10 := []int{127, -2147483648, 128}
            _ = v10
        }
    } else {
        if (-128 <= -56) {
            var v11 int16 = -1
            _ = v11
        }
        v12 := []int{9223372036854775807, -163, -2147483648}
        _ = v12
    }
    p0 = p0
    _ = p0
    _ = v6
    return 90
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 1.5
    var v3 float64 = v2
    var v13 uint64 = 51
    var v14 int16 = -486
    v15 := GenS1{F0: 767, F1: 1, F2: ("hello" + (((("\n\t" + ("hello" + "")) + (("a" + "hello") + "\n\t")) + ("a" + "foo bar")) + ("foo bar" + "a")))}
    delete(v1, 0)
    v16 := []int{}
    switch ((("hello" + "hello") + "foo bar") + ("\n\t" + ("\n\t" + "a"))) {
    case "a":
        v17 := GenS2{F0: 42, F1: 547}
        _ = v17
    case "\n\t":
        var v18 int16 = 368
        _ = v18
    }
    var v19 uint8 = 127
    v3 = v3
    v20 := GenS1{F0: -394, F1: 42, F2: "hello"}
    var v21 int16 = 1000
    v22 := v1[4294967295]
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            v14 -= -1
            _ = fi24
        }
        var v25 uint = 4294967295
        _ = fi23
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
