package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0(p0 int, p1 int8) uint {
    if (4294967295 <= 42) {
        if (42 <= 0) {
            p1 += p1
            v6 := GenS1{F0: 1, F1: 1}
            _ = v6
        } else {
            v7 := []int{777, (p0 * 9223372036854775807), (-(p0))}
            var v8 float32 = 1.5
            var v9 int8 = (p1 >> ((42) & 63))
            _ = v7
            _ = v8
            _ = v9
        }
    }
    p1 += p1
    v10 := GenS1{F0: 1, F1: 0}
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint64 = 0
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return 65535
}
func genFn1(p0 int, p1 string, p2 bool) int {
    switch ((((p1 + "") + ("hello" + ((p1 + ((p1 + p1) + p1)) + ((("" + p1) + p1) + ("hello" + ((p1 + p1) + ("foo bar" + (p1 + (p1 + p1))))))))) + "foo bar") + p1) {
    case "a":
        fmt.Println(p0)
    case "foo bar":
        var v13 float32 = (0.0 * -2.25)
        _ = v13
    }
    p1 = "a"
    fmt.Println(p1, p2)
    switch (p1 + "\n\t") {
    case "hello":
        switch ((("" + ("a" + ((p1 + ((p1 + ("\n\t" + ((p1 + p1) + (p1 + ((("" + (p1 + ((p1 + p1) + (p1 + "\n\t")))) + p1) + p1))))) + (p1 + p1))) + p1))) + p1) + p1) {
        case "a":
            v14 := []int{p0}
            _ = v14
        case "\n\t":
            p0 = (p0 >> ((481) & 63))
        case "foo bar":
            if false {
                v15 := GenS1{F0: 1, F1: 42}
                var v16 int8 = 42
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 3; fi17++ {
                    if (genFn0((^(fi17)), 42) > 0) {
                        var v18 int8 = 42
                        _ = v18
                    }
                    p1 = (p1 + "foo bar")
                    _ = fi17
                }
            }
        default:
            for fi19 := 0; fi19 < 20; fi19++ {
                v20 := GenS1{F0: 9223372036854775807, F1: 1}
                p2 = ((!p2) || ((fi19 ^ fi19) != fi19))
                _ = fi19
                _ = v20
            }
        }
    case "foo bar":
        var v21 bool = false
        _ = v21
    }
    switch (p0 + p0) {
    case -267:
        if true {
            p2 = (-19 <= -128)
        } else {
            switch p1 {
            case "foo bar":
                var v22 uint32 = 2147483647
                _ = v22
            }
            p0 = 42
        }
    case 4294967295:
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 float32 = 3.14159
    var v2 string = ""
    var v3 int64 = -131
    v4 := make(map[int]uint32)
    var v5 uint8 = 0
    v5 = 0
    var v23 uint32 = 2147483647
    var v24 uint64 = 4294967295
    if (true || true) {
        for fi25 := 0; fi25 < 1; fi25++ {
            v26 := GenS1{F0: v3, F1: int16(v5)}
            _ = fi25
            _ = v26
        }
        var v27 uint32 = v23
        _ = v27
    }
    var v28 uint8 = (^(v5))
    v24 = (^(v24))
    var v29 bool = true
    var v30 int8 = 100
    fmt.Println(v23, v1)
    for fi31 := 0; fi31 < 19; fi31++ {
        for fi32 := 0; fi32 < 10; fi32++ {
            _ = len(v4)
            v3 = (-(v3))
            _ = fi32
        }
        v33 := GenS1{F0: (v3 * v3), F1: 42}
        _ = fi31
        _ = v33
    }
    fmt.Println(v24, v30, v30)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
