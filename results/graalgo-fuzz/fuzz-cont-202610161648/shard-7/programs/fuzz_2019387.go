package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    v4 := []string{"hello", "hello", (("" + "foo bar") + "a")}
    switch "" {
    case "foo bar":
        if ((882 == 2) && true) {
            var v5 uint64 = 18446744073709551615
            if false {
                v6 := GenS1{F0: 0, F1: v5}
                if true {
                    v7 := GenS1{F0: 1, F1: v5}
                    v8 := GenS2{F0: true}
                    _ = v7
                    _ = v8
                }
                _ = v6
            }
            var v9 uint = 255
            _ = v5
            _ = v9
        } else {
            var v10 int8 = 61
            fmt.Println(v10)
            var v11 uint32 = 65535
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint16 = 255
        _ = fi12
        _ = v13
    }
    if (-201 > 32768) {
        var v14 uint64 = 42
        for fi15 := 0; fi15 < 15; fi15++ {
            if (("hello" + "") <= "foo bar") {
                var v16 int8 = 127
                v14 -= v14
                _ = v16
            }
            v14 -= v14
            v4 = append(v4, "a")
            _ = fi15
        }
        v14 = v14
        _ = v14
    }
    var v17 int = 32768
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 int = (v17 * fi18)
        if (190 > 4294967295) {
            for fi20 := 0; fi20 < 11; fi20++ {
                v19 = (^(v17))
                var v21 uint64 = 4294967295
                _ = fi20
                _ = v21
            }
        } else {
            fmt.Println(fi18, v17)
            v17 += 0
            var v22 int = v17
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        v4 = append(v4, ("hello" + "foo bar"))
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 float64 = (0.1 + 1e10)
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    v17 -= v17
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        switch "a" {
        case "\n\t":
            v17 = v17
        }
        if true {
            v17 -= (fi26 | (fi26 << ((595) & 63)))
            for fi27 := 0; fi27 < 4; fi27++ {
                switch ((("\n\t" + "\n\t") + (("hello" + "foo bar") + (("a" + ("" + "\n\t")) + "hello"))) + ("foo bar" + ("foo bar" + "a"))) {
                case "":
                    v28 := GenS3{F0: 42}
                    _ = v28
                default:
                    v29 := GenS1{F0: 841, F1: 1}
                    _ = v29
                }
                _ = fi27
            }
        }
        v17 += (fi26 + 616)
        _ = fi26
    }
    fmt.Println(v17)
    for fi30 := 0; fi30 < 4; fi30++ {
        for fi31 := 0; fi31 < 10; fi31++ {
            if false {
                fmt.Println(fi31)
                var v32 uint32 = 680
                fmt.Println(fi31)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v33 := GenS2{F0: false}
                _ = v32
                _ = v33
            } else {
                v17 += fi31
                switch (v17 & 1) {
                case 986:
                    var v34 uint16 = 44
                    _ = v34
                default:
                    v17 -= -329
                }
            }
            if false {
                if len(v4) > 0 {
                    v4[0] = ((("foo bar" + ((("\n\t" + ("foo bar" + ("hello" + "hello"))) + "") + (("hello" + "foo bar") + "a"))) + "a") + "")
                }
            }
            _ = fi31
        }
        var v35 uint8 = 128
        v36 := GenS4{F0: -62, F1: 0}
        _ = fi30
        _ = v35
        _ = v36
    }
    var v37 uint8 = 128
    var v38 float32 = 0.1
    switch "hello" {
    case "":
        v39 := GenS3{F0: -9223372036854775808}
        _ = v39
    }
    switch v17 {
    case 128:
        v40 := GenS3{F0: 347}
        _ = v40
    case 42:
        fmt.Println(v17, v37)
    }
    v4 = append(v4, (("hello" + "\n\t") + "a"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
