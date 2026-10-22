package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int32, p1 int, p2 int64) uint64 {
    if false {
        switch (("\n\t" + "a") + "a") {
        case "a":
            var v5 uint8 = 240
            _ = v5
        case "foo bar":
            p2 = (p2 + p2)
        }
        switch p1 {
        case 128:
            fmt.Println(p2, p0)
        case 0:
            var v6 uint8 = 101
            _ = v6
        case 9223372036854775807:
            p1 -= p1
        }
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int = 0
        switch (-(v8)) {
        case 32768:
            if false {
                p1 += 32768
                var v9 int = p1
                var v10 uint = 675
                _ = v10
                _ = v9
            }
        case 127:
            p2 = (p2 + p2)
        case 767:
            fmt.Println(p1, p1, p0)
        }
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: 32767}
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 65535
    v14 := GenS1{F0: 1000}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 2
}

func main() {
    var v1 uint16 = 65535
    var v2 float64 = float64(float32(-2.25))
    v3 := make(map[string]int)
    var v4 float32 = (1.5 - 1.5)
    if false {
        if true {
            v15 := GenS1{F0: -915}
            var v16 int64 = 42
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 13; fi17++ {
                v18 := v3[("a" + "foo bar")]
                var v19 string = ""
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        v20 := GenS2{F0: 32767}
        v3[("a" + "foo bar")] = 32768
        _ = v20
    } else {
        fmt.Println(v1, v2, v1)
        if (!false) {
            v4 = v4
            switch "\n\t" {
            case "a":
                v21 := GenS1{F0: 42}
                _ = v21
            case "\n\t":
                v4 = ((0.0 / 1.5) + 3.14159)
            }
            v22 := GenS2{F0: 455}
            _ = v22
        }
        fmt.Println(v2, v1)
    }
    if true {
        v23 := GenS1{F0: -32768}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 6; fi24++ {
            fmt.Println(fi24, v4, v4)
            delete(v3, "hello")
            v1 = (v1 << ((65535) & 63))
            _ = fi24
        }
        switch -822 {
        case 127:
            var v25 string = ("foo bar" + "foo bar")
            _ = v25
        case 9223372036854775807:
            var v26 string = ((("a" + "a") + ("foo bar" + (((("" + ((("hello" + ("a" + "foo bar")) + (("foo bar" + "") + (((("a" + "") + "") + "a") + "a"))) + ("\n\t" + "foo bar"))) + "foo bar") + "\n\t") + "hello"))) + "hello")
            _ = v26
        }
        v27 := GenS3{F0: v1}
        _ = v27
    }
    delete(v3, "\n\t")
    var v28 int32 = 2147483647
    v2 = (float64(v4) / (v2 / v2))
    fmt.Println(v4, v2, v28)
    v29 := GenS3{F0: 1}
    v4 = (v4 / -2.25)
    v30 := GenS1{F0: 780}
    _ = len(v3)
    if false {
        var v31 int64 = 42
        for fi32 := 0; fi32 < 10; fi32++ {
            v4 += (float32(v2) + float32(v2))
            _ = fi32
        }
        _ = v31
    } else {
        v33 := GenS4{F0: 575}
        _ = v33
    }
    var v34 uint = 1
    if true {
        v35 := []int{42, 0, 1}
        var v36 int8 = 100
        if true {
            switch -115 {
            case 669:
                v36 = (v36 & -1)
            }
        }
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
}
