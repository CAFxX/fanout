package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = float32(0.0)
    var v3 uint = 255
    var v4 uint64 = 370
    delete(v1, 582)
    if false {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 int = fi5
            if (-435 > -32768) {
                v7 := GenS1{F0: 1}
                var v8 int8 = 1
                v9 := GenS2{F0: v8, F1: "a"}
                _ = v7
                _ = v8
                _ = v9
            } else {
                var v10 uint64 = (v4 * v4)
                _ = v10
            }
            _ = fi5
            _ = v6
        }
    } else {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 bool = true
            if (!(!true)) {
                if (238 >= 65535) {
                    v13 := []int{(fi11 % (((fi11 | (fi11 + (-(fi11))))) | 1))}
                    v12 = ((0 < 127) && (2147483647 < int32(fi11)))
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = v13
                }
                v14 := v1[(^(fi11))]
                _ = v14
            }
            fmt.Println(fi11)
            _ = fi11
            _ = v12
        }
        v15 := GenS3{F0: 4294967295}
        _ = v15
    }
    var v16 bool = true
    v17 := GenS4{F0: (42 < int16(v4))}
    for fi18 := 0; fi18 < 14; fi18++ {
        switch (("hello" + "foo bar") + "hello") {
        case "foo bar":
            if false {
                switch fi18 {
                case 9223372036854775807:
                    fmt.Println(fi18)
                case 0:
                    v2 += (v2 + 0.1)
                }
                switch ("hello" + "foo bar") {
                case "a":
                    v19 := GenS2{F0: -128, F1: ((("\n\t" + "\n\t") + "foo bar") + (("a" + "foo bar") + "a"))}
                    _ = v19
                case "foo bar":
                    for fi20 := 0; fi20 < 1; fi20++ {
                        v4 -= (v4 + 4294967295)
                        var v21 int8 = -11
                        _ = fi20
                        _ = v21
                    }
                }
                var v22 uint32 = 65535
                _ = v22
            }
        case "a":
            switch "\n\t" {
            case "foo bar":
                v23 := []int{int(v3), 128, fi18}
                _ = v23
            }
        }
        v24 := GenS2{F0: 127, F1: ""}
        var v25 int32 = 143
        _ = fi18
        _ = v24
        _ = v25
    }
    v26 := GenS4{F0: v16}
    v27 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
