package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("foo bar" + "\n\t"), "\n\t", "foo bar", ""}
    v3 := make(map[int]bool)
    var v4 int8 = 42
    v5 := []string{(("a" + "\n\t") + ""), "hello", ""}
    switch "hello" {
    case "":
        v4 += v4
    case "\n\t":
        for fi6 := 0; fi6 < 17; fi6++ {
            fmt.Println(v4)
            _ = fi6
        }
    case "a":
        v4 = (v4 - v4)
    default:
        switch ((("foo bar" + "") + ("" + "\n\t")) + "\n\t") {
        case "foo bar":
            v4 -= v4
        case "":
            var v7 int = 128
            _ = v7
        case "hello":
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 int32 = -880
                var v10 int64 = -1
                for fi11 := 0; fi11 < 13; fi11++ {
                    v10 -= 0
                    _ = fi11
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        }
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v4 += v4
        if ((128 == 42) || false) {
            switch (fi12 >> ((65535) & 63)) {
            case 32767:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case -235:
                v4 += (v4 + v4)
            case -9223372036854775808:
                v4 = 127
            }
        } else {
            var v13 uint32 = 42
            var v14 string = ("hello" + "\n\t")
            _ = v13
            _ = v14
        }
        if (!(false && true)) {
            switch fi12 {
            case 127:
                v15 := GenS1{F0: 414, F1: 0}
                _ = v15
            }
            var v16 uint16 = 640
            for fi17 := 0; fi17 < 20; fi17++ {
                v16 -= 255
                v16 = (^(v16))
                v4 = v4
                _ = fi17
            }
            _ = v16
        } else {
            v18 := GenS2{F0: "\n\t", F1: -353, F2: true}
            v2 = append(v2, "a")
            _ = v18
        }
        _ = fi12
    }
    v19 := GenS1{F0: 514, F1: 142}
    v4 += (v4 >> ((736) & 63))
    fmt.Println(v4)
    fmt.Println(v4)
    fmt.Println(v4)
    var v20 int8 = (v4 + (v4 | (v4 | -1)))
    _ = len(v2)
    v4 -= v4
    v3[-1] = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
