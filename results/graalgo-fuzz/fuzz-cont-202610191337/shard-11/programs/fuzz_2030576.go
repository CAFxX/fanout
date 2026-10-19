package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint16) uint8 {
    var v5 int16 = 842
    v6 := []int{128, 1}
    fmt.Println(p0, p0)
    _ = p0
    _ = v5
    _ = v6
    return 127
}
func genFn1(p0 int8, p1 uint, p2 uint8) int64 {
    var v7 uint = (p1 | (p1 ^ (p1 | p1)))
    v8 := GenS1{F0: "a"}
    v9 := GenS1{F0: (("" + ("a" + ((((("a" + "foo bar") + "hello") + (("foo bar" + ("a" + (("foo bar" + "") + ((("" + "foo bar") + "") + "a")))) + ((((("\n\t" + "a") + ("\n\t" + ("\n\t" + ("" + ("a" + ""))))) + "hello") + "hello") + "\n\t"))) + "") + "foo bar"))) + ("hello" + ("hello" + "")))}
    v10 := GenS1{F0: "hello"}
    v11 := GenS1{F0: ""}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := []string{"\n\t", "hello"}
    var v2 int64 = 0
    var v3 int64 = (v2 & 1)
    v4 := []int{128, 0, -1, 9223372036854775807}
    for fi12 := 0; fi12 < 3; fi12++ {
        switch fi12 {
        case -948:
            fmt.Println(v3, v3, v2)
        case -432:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        v13 := GenS1{F0: ("" + "hello")}
        _ = fi12
        _ = v13
    }
    v3 = (v3 << ((65535) & 63))
    for fi14 := 0; fi14 < 1; fi14++ {
        if (!(false || (((845 == 1) && (750 > 255)) || false))) {
            for fi15 := 0; fi15 < 13; fi15++ {
                v2 = (-(v2))
                _ = fi15
            }
        } else {
            v2 = v2
            var v16 int32 = 42
            for fi17 := 0; fi17 < 5; fi17++ {
                v2 = genFn1(42, 143, 13)
                var v18 int8 = 0
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        v19 := GenS1{F0: "foo bar"}
        var v20 float32 = 0.1
        _ = fi14
        _ = v19
        _ = v20
    }
    v1 = append(v1, "a")
    v21 := GenS1{F0: ("a" + "hello")}
    v3 += v3
    v1 = append(v1, "\n\t")
    var v22 bool = true
    var v23 uint32 = 2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
