package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 string; F1 string; F2 uint64 }

func main() {
    var v1 float32 = 0.0
    var v2 int32 = 42
    v3 := make(map[int]int)
    var v4 int8 = -26
    v5 := make(map[int]bool)
    switch ("a" + ("a" + ("a" + "\n\t"))) {
    case "a":
        v6 := GenS1{F0: 127}
        _ = v6
    }
    fmt.Println(v4, v4)
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 int = 666
            if (true || (!(-9223372036854775808 != 1))) {
                var v10 uint64 = 1
                switch ((("hello" + "foo bar") + "") + "hello") {
                case "a":
                    var v11 int64 = 1
                    _ = v11
                case "foo bar":
                    var v12 float64 = ((1e10 / 0.1) + 1e10)
                    _ = v12
                }
                fmt.Println(v10, fi7, v10)
                _ = v10
            } else {
                var v13 int8 = (v4 * -30)
                v9 += (fi7 >> ((65) & 63))
                _ = v13
            }
            v14 := GenS2{F0: 839, F1: (fi8 - (v9 >> ((65535) & 63))), F2: 0}
            _ = fi8
            _ = v14
            _ = v9
        }
        _ = fi7
    }
    var v15 uint = 1
    v16 := GenS3{F0: ((((("foo bar" + ("hello" + "")) + "") + "") + "\n\t") + ""), F1: (("\n\t" + "foo bar") + "\n\t"), F2: 746}
    v17 := GenS3{F0: "hello", F1: "hello", F2: 0}
    if true {
        v2 = int32(v4)
    } else {
        fmt.Println(v2)
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 string = ("a" + "")
            _ = fi18
            _ = v19
        }
    }
    v20 := GenS3{F0: ("foo bar" + "a"), F1: "a", F2: 833}
    v21 := GenS2{F0: 0, F1: -91, F2: 543}
    switch ("\n\t" + "hello") {
    case "foo bar":
        if (true && (!((42 != 128) || true))) {
            for fi22 := 0; fi22 < 18; fi22++ {
                for fi23 := 0; fi23 < 4; fi23++ {
                    if false {
                        v4 += v4
                    }
                    var v24 uint32 = 0
                    v2 -= v2
                    v2 = v2
                    _ = fi23
                    _ = v24
                }
                _ = fi22
            }
        }
    case "a":
        var v25 uint32 = 4294967295
        _ = v25
    }
    v26 := GenS3{F0: "", F1: "foo bar", F2: 2}
    var v27 float64 = 1e10
    v2 = (v2 >> ((v15) & 63))
    for fi28 := 0; fi28 < 9; fi28++ {
        v15 = v15
        v27 = (3.14159 + v27)
        _ = fi28
    }
    var v29 uint64 = 137
    var v30 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
