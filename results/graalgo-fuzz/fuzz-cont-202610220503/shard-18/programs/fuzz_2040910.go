package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64; F1 int; F2 int }
type GenS3 struct { F0 uint; F1 uint8 }

func main() {
    v1 := []string{}
    v2 := make(map[int]bool)
    v3 := []string{"\n\t", ((("" + "") + "foo bar") + ("\n\t" + ((("foo bar" + ("hello" + "hello")) + "\n\t") + ("foo bar" + ("\n\t" + (("a" + "") + ("" + "hello"))))))), ((("\n\t" + "hello") + ("" + "\n\t")) + ("foo bar" + ("a" + ("hello" + "a")))), "hello"}
    var v4 int64 = 873
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: -281, F1: -270}
    if true {
        for fi7 := 0; fi7 < 9; fi7++ {
            _ = len(v1)
            var v8 float32 = -2.25
            _ = fi7
            _ = v8
        }
    }
    v4 = (v4 / (((v4 * v4)) | 1))
    v9 := GenS1{F0: 42, F1: 155}
    for fi10 := 0; fi10 < 2; fi10++ {
        _ = len(v1)
        var v11 int16 = 42
        _ = fi10
        _ = v11
    }
    fmt.Println(v4)
    v4 -= v4
    switch "a" {
    case "a":
        switch "hello" {
        case "a":
            var v12 int64 = v4
            _ = v12
        case "\n\t":
            var v13 int8 = 9
            _ = v13
        case "":
            v4 += (v4 ^ v4)
        }
    case "\n\t":
        var v14 int = 127
        _ = v14
    case "":
        v4 = v4
    default:
        var v15 string = "hello"
        _ = v15
    }
    switch ("\n\t" + "") {
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            v4 = v4
        case "":
            var v16 int32 = 637
            _ = v16
        case "a":
            v17 := GenS1{F0: -2147483648, F1: 32767}
            _ = v17
        default:
            for fi18 := 0; fi18 < 17; fi18++ {
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 float32 = float32(1e10)
                    v21 := GenS2{F0: 2, F1: 2147483647, F2: 200}
                    v20 = v20
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                fmt.Println(fi18)
                _ = len(v3)
                _ = fi18
            }
        }
    }
    if true {
        v4 += (^(v4))
        switch ("a" + "\n\t") {
        case "\n\t":
            fmt.Println(v4)
        case "":
            var v22 int8 = 100
            _ = v22
        case "hello":
            var v23 uint = 2
            _ = v23
        }
    } else {
        for fi24 := 0; fi24 < 2; fi24++ {
            v4 = 0
            _ = fi24
        }
        switch 32768 {
        case 127:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 128:
            var v25 int = 127
            _ = v25
        }
        switch 9223372036854775807 {
        case 2147483647:
            if ((42 >= 2147483647) && true) {
                var v26 int16 = -872
                var v27 float32 = 1.5
                v28 := GenS1{F0: 2147483647, F1: (v26 % ((v26) | 1))}
                _ = v26
                _ = v27
                _ = v28
            } else {
                v4 -= v4
            }
        case 1:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 37:
            v4 -= (v4 << ((6) & 63))
        }
    }
    v29 := GenS3{F0: 385, F1: 227}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
