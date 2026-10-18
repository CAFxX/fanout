package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int; F1 uint64; F2 uint32 }
func genFn0(p0 int32, p1 int32, p2 bool) string {
    if p2 {
        switch "" {
        case "":
            fmt.Println(p1, p1)
        case "foo bar":
            p0 = 44
        }
        v6 := GenS1{F0: ("\n\t" + (((("" + "") + "a") + "hello") + ("hello" + "a")))}
        p0 = 42
        _ = v6
    } else {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 float64 = 3.14159
            _ = fi7
            _ = v8
        }
        v9 := []int{}
        var v10 string = (("a" + "a") + ("" + ("\n\t" + "\n\t")))
        _ = v10
        _ = v9
    }
    switch 2147483647 {
    case 9223372036854775807:
        for fi11 := 0; fi11 < 1; fi11++ {
            for fi12 := 0; fi12 < 11; fi12++ {
                if (-57 != 42) {
                    var v13 int16 = 479
                    _ = v13
                } else {
                    fmt.Println(p1, p0, p1)
                }
                p1 = (p0 | p1)
                v14 := GenS2{F0: (p0 & p1), F1: 9223372036854775807}
                p1 = (^(p0))
                _ = fi12
                _ = v14
            }
            switch (fi11 << ((65535) & 63)) {
            case 127:
                fmt.Println(p1, p2, fi11)
            default:
                for fi15 := 0; fi15 < 10; fi15++ {
                    v16 := GenS1{F0: ""}
                    p2 = false
                    _ = fi15
                    _ = v16
                }
            }
            var v17 float32 = 1.5
            _ = fi11
            _ = v17
        }
    case -1:
        var v18 uint64 = 0
        _ = v18
    case 1:
        for fi19 := 0; fi19 < 10; fi19++ {
            fmt.Println(p1)
            fmt.Println(p1, p0)
            var v20 float64 = float64((0.1 - (0.1 - (1e10 * 0.0))))
            _ = fi19
            _ = v20
        }
    default:
        p1 = p1
    }
    var v21 float32 = (1.5 * (-2.25 - 1.5))
    var v22 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return "a"
}

func main() {
    var v1 uint32 = 828
    var v2 int = -1
    var v3 bool = false
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]bool)
    fmt.Println(v1, v1)
    v23 := []int{v2, (v2 << ((42) & 63)), (^(v2))}
    v2 = (^(v2))
    switch "\n\t" {
    case "hello":
        var v24 int32 = 0
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        v26 := GenS1{F0: ("\n\t" + "foo bar")}
        var v27 float32 = -2.25
        _ = fi25
        _ = v26
        _ = v27
    }
    switch ((("" + ("hello" + ("" + "a"))) + "hello") + (("hello" + "hello") + "foo bar")) {
    case "":
        for fi28 := 0; fi28 < 19; fi28++ {
            fmt.Println(fi28, fi28, v2)
            _ = fi28
        }
    case "hello":
        v29 := GenS1{F0: ("\n\t" + "a")}
        _ = v29
    case "a":
        v4[v1] = ((-(v2)) <= v2)
    default:
        for fi30 := 0; fi30 < 20; fi30++ {
            fmt.Println(v3)
            v31 := GenS3{F0: 0}
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 uint64 = 4294967295
                switch "hello" {
                case "\n\t":
                    v3 = ((v3 && v3) || false)
                case "a":
                    v34 := GenS4{F0: 1}
                    _ = v34
                case "":
                    fmt.Println(v2, v3)
                }
                _ = fi32
                _ = v33
            }
            _ = fi30
            _ = v31
        }
    }
    if false {
        var v35 uint32 = v1
        _ = v35
    }
    switch "hello" {
    case "\n\t":
        for fi36 := 0; fi36 < 19; fi36++ {
            var v37 bool = true
            _ = fi36
            _ = v37
        }
    case "a":
        v1 = v1
    case "hello":
        fmt.Println(v3)
    default:
        var v38 int16 = -32768
        _ = v38
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        switch (("a" + (("foo bar" + "\n\t") + "hello")) + (("a" + "a") + "foo bar")) {
        case "":
            var v40 int64 = 42
            _ = v40
        case "a":
            for fi41 := 0; fi41 < 4; fi41++ {
                var v42 uint32 = 42
                v23 = append(v23, (v2 & (-(fi39))))
                _ = fi41
                _ = v42
            }
        case "\n\t":
            var v43 uint32 = v1
            _ = v43
        default:
            var v44 float32 = ((float32(3.14159) + 0.0) * -2.25)
            _ = v44
        }
        _ = fi39
    }
    v45 := GenS5{F0: 324, F1: -432}
    v46 := v5[42]
    v47 := GenS6{F0: 128, F1: 2, F2: (v1 | v1)}
    v2 = (v2 + (-(v2)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
