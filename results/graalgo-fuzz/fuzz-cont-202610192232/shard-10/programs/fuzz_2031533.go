package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint16; F2 bool }

func main() {
    var v1 int16 = 1000
    v2 := []int{1}
    v3 := make(map[int]int)
    v4 := []string{}
    switch "hello" {
    case "foo bar":
        v5 := GenS1{F0: "foo bar", F1: 1, F2: -2147483648}
        _ = v5
    case "a":
        v3[9223372036854775807] = 0
    }
    var v6 uint32 = 952
    var v7 int8 = 127
    v8 := GenS1{F0: "\n\t", F1: 42, F2: 42}
    var v9 bool = ((127 >= 128) && ((v7 <= v7) && (("\n\t" + "\n\t") == "hello")))
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            v6 = v6
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v9 = (42 != fi11)
            _ = fi11
        }
        if false {
            v12 := GenS1{F0: ("a" + "\n\t"), F1: 255, F2: -747}
            _ = v12
        } else {
            v6 += 42
            v13 := GenS2{F0: 1}
            _ = v13
        }
        _ = fi10
    }
    v9 = (255 != 433)
    v3[2147483647] = int(v1)
    v1 = v1
    if v9 {
        v1 = 0
        if (!true) {
            v7 = 127
        } else {
            for fi14 := 0; fi14 < 15; fi14++ {
                v9 = false
                var v15 uint32 = 0
                _ = fi14
                _ = v15
            }
            fmt.Println(v6, v9, v7)
        }
    } else {
        var v16 uint8 = 1
        for fi17 := 0; fi17 < 4; fi17++ {
            v1 += (v1 & v1)
            v16 += 201
            _ = fi17
        }
        if len(v4) > 0 {
            v4[0] = ""
        }
        _ = v16
    }
    fmt.Println(v6)
    fmt.Println(v7)
    fmt.Println(v1)
    v7 += (^(v7))
    switch 861 {
    case 2147483647:
        v1 = -1
    default:
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS3{F0: 2, F1: 52, F2: v9}
            switch "a" {
            case "hello":
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 uint16 = 32767
                    fmt.Println(v7, v1, fi18)
                    var v22 int32 = -1
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            v9 = v9
            _ = fi18
            _ = v19
        }
    }
    v9 = (!true)
    var v23 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
