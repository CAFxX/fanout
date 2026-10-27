package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    if false {
        switch "a" {
        case "hello":
            v4 := GenS1{F0: -9223372036854775808, F1: -53}
            _ = v4
        case "\n\t":
            var v5 int8 = -56
            _ = v5
        }
        var v6 uint64 = 484
        _ = v6
    }
    var v7 int32 = 42
    var v8 float32 = 1.5
    var v9 uint = 1
    switch 0 {
    case -231:
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(v8, fi10)
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int = fi11
                fmt.Println(v12, v12, fi10)
                v12 += 32768
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    case -2147483648:
        v8 += v8
    case 127:
        for fi13 := 0; fi13 < 12; fi13++ {
            switch 1 {
            case -920:
                v7 = v7
            }
            _ = fi13
        }
    }
    if true {
        for fi14 := 0; fi14 < 4; fi14++ {
            switch "a" {
            case "a":
                fmt.Println(v9)
            case "\n\t":
                v15 := []int{fi14, fi14, fi14}
                _ = v15
            case "foo bar":
                v16 := GenS1{F0: fi14, F1: -90}
                _ = v16
            }
            var v17 float32 = (v8 * v8)
            _ = fi14
            _ = v17
        }
        switch -927 {
        case 556:
            switch (("foo bar" + "") + (("" + "hello") + (("" + "a") + ("a" + (("hello" + "\n\t") + ("\n\t" + ("" + ("foo bar" + "\n\t")))))))) {
            case "foo bar":
                v7 += v7
            case "\n\t":
                v18 := []int{}
                _ = v18
            case "":
                v19 := GenS1{F0: 32767, F1: -128}
                _ = v19
            }
        default:
            if false {
                v9 = 42
            }
        }
    }
    v8 = 0.0
    var v20 uint8 = 132
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v7)
    case "\n\t":
        v7 = 0
    case "hello":
        for fi21 := 0; fi21 < 14; fi21++ {
            v22 := []int{(fi21 ^ (fi21 ^ 4294967295)), (fi21 * (fi21 & fi21)), fi21}
            _ = fi21
            _ = v22
        }
    }
    switch (((("foo bar" + "a") + "foo bar") + "") + ("" + (("\n\t" + "\n\t") + "\n\t"))) {
    case "\n\t":
        v7 += (-(v7))
    case "foo bar":
        var v23 float32 = v8
        _ = v23
    case "hello":
        fmt.Println(v7)
    }
    var v24 int32 = 42
    v25 := GenS1{F0: 256, F1: -121}
    switch ((("\n\t" + "a") + "hello") + "") {
    case "":
        var v26 uint8 = v20
        _ = v26
    case "a":
        v27 := GenS1{F0: 128, F1: 121}
        _ = v27
    }
    delete(v1, ("\n\t" + ""))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
