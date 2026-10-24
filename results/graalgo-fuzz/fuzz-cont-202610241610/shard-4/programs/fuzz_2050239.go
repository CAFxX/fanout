package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 string; F1 uint16; F2 int }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    var v4 int = 128
    v5 := make(map[string]bool)
    var v6 bool = false
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(fi7, v6, v6)
        var v8 int8 = 0
        if false {
            v9 := GenS1{F0: 126, F1: 42}
            fmt.Println(v8, v6)
            if true {
                v6 = v6
                for fi10 := 0; fi10 < 20; fi10++ {
                    fmt.Println(v6)
                    v4 = (fi7 << ((512) & 63))
                    switch "\n\t" {
                    case "\n\t":
                        v11 := GenS2{F0: "a", F1: 67, F2: (^(fi7))}
                        _ = v11
                    }
                    _ = fi10
                }
                fmt.Println(v8)
            } else {
                v8 = (v8 + -1)
                if true {
                    var v12 uint16 = 1
                    _ = v12
                }
                var v13 string = "foo bar"
                _ = v13
            }
            _ = v9
        } else {
            v4 -= v4
        }
        _ = fi7
        _ = v8
    }
    delete(v3, ("" + "a"))
    var v14 int = -1000
    var v15 string = ("hello" + ("hello" + ((("" + "a") + "") + ("a" + "hello"))))
    fmt.Println(v14, v15)
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            fmt.Println(v15, v14, v6)
            v18 := []int{}
            fmt.Println(fi16)
            _ = fi17
            _ = v18
        }
        var v19 int16 = 1000
        switch (v15 + v15) {
        case "hello":
            var v20 int64 = 42
            _ = v20
        case "\n\t":
            v15 = v15
        case "foo bar":
            v21 := GenS2{F0: v15, F1: 255, F2: v4}
            _ = v21
        }
        _ = fi16
        _ = v19
    }
    fmt.Println(v6, v6)
    v22 := GenS3{F0: 18446744073709551615, F1: 1}
    for fi23 := 0; fi23 < 7; fi23++ {
        v4 -= 127
        for fi24 := 0; fi24 < 17; fi24++ {
            if (42 <= -1) {
                v25 := []int{(-(fi23)), 795, (v4 * 32768)}
                v26 := GenS2{F0: ("" + ("\n\t" + v15)), F1: 65535, F2: 127}
                _ = v25
                _ = v26
            }
            _ = fi24
        }
        if true {
            if (0.1 <= ((float32(0.1) - 1.5) * 3.14159)) {
                v27 := []int{}
                v28 := GenS2{F0: v15, F1: 42, F2: 2147483647}
                var v29 int = (fi23 / (((v4 << ((4294967295) & 63))) | 1))
                _ = v27
                _ = v28
                _ = v29
            } else {
                fmt.Println(fi23, v6)
                v30 := GenS4{F0: 592}
                v6 = v6
                _ = v30
            }
            var v31 uint16 = 915
            _ = v31
        } else {
            var v32 int64 = 9223372036854775807
            _ = v32
        }
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
